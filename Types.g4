grammar Types;

ATOMIC_INT32: 'int32';
ATOMIC_UINT32: 'uint32';
ATOMIC_INT64: 'int64';
ATOMIC_UINT64: 'uint64';
ATOMIC_STRING: 'string';
ATOMIC_BOOL: 'bool';
ATOMIC_BYTES: 'bytes';
ATOMIC_DOUBLE: 'double';
ATOMIC_FLOAT: 'float';
ATOMIC_CHAR: 'char';

IDENT: [a-zA-Z_][A-Za-z0-9_]*;
FIXED_FIELDNAME: '"' IDENT '"';
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: value = type EOF;

type:
	'[' elemType = type ']'							# List
	| name = IDENT									# TypeVar
	| '$' name = IDENT 						        # PackVar
	| '(' types = tupleTypeList ')'					# Tuple
	| '|' types = unionTypeList '|'					# Union
	| '{' fields = fieldList '}'					# Struct
	| name = IDENT '<' types = tupleTypeList '>'	# Named
	| atom = atomic									# AtomicType;

tupleTypeList:
	pattern = type '...'						# TupleTypeListOnlyExpansion
	| (types += type ',')+ finalType = type		# TupleTypeListNoExpansion
	| (types += type ',')+ pattern = type '...'	# TupleTypeListExpansion
	|											# TupleTypeListEmpty;

unionTypeList:
	pattern = type '...'						# UnionTypeListOnlyExpansion
	| (types += type '|')+ finalType = type		# UnionTypeListNoExpansion
	| (types += type '|')+ pattern = type '...'	# UnionTypeListExpansion
	|											# UnionTypeListEmpty;

fieldList:
	'$' namePattern = IDENT ':' typePattern = type '...'									            # FieldListOnlyExpansion
	| (names += fieldName ':' types += type ',')+ finalName = fieldName ':' finalType = type	        # FieldListNoExpansion
	| (names += fieldName ':' types += type ',')+ '$' namePattern = IDENT ':' typePattern = type '...'	# FieldListExpansion
	|				                                                                                    # FieldListEmpty;

fieldName:
	name = FIXED_FIELDNAME	# FieldNameFixed
	| name = IDENT			# FieldNameVariable;

atomic:
	name = ATOMIC_INT32
	| name = ATOMIC_UINT32
	| name = ATOMIC_INT64
	| name = ATOMIC_UINT64
	| name = ATOMIC_STRING
	| name = ATOMIC_BOOL
	| name = ATOMIC_BYTES
	| name = ATOMIC_DOUBLE
	| name = ATOMIC_FLOAT
	| name = ATOMIC_CHAR;