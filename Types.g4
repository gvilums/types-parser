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

// root rule
start: value = type EOF;

type:
	t = regularType			# TypeNonUnion
	| types = unionTypeList	# TypeUnion;

regularType:
	'[' elemType = type ']'							# List
	| name = IDENT									# TypeVar
	| '$' name = IDENT								# PackVar
	| '{' fields = fieldList '}'					# Struct
	| atom = atomic									# AtomicType
	| name = IDENT '<' elemType = type '>'			# NamedSingle
	| name = IDENT '<' types = tupleTypeList '>'	# Named
	| '(' types = tupleTypeList ')'					# Tuple
	| '(' innerType = type ')'						# Parenthesized;

tupleTypeList:
	(types += type ',')* (pattern = type '...')	# TupleTypeListExpansion
	| (types += type ',')+ (types += type)?		# TupleTypeListNoExpansion
	|											# TupleTypeListEmpty;

unionTypeList:
	(types += regularType '|')* (pattern = regularType '...')	# UnionTypeListExpansion
	| (types += regularType '|')+ (types += regularType)?		# UnionTypeListNoExpansion;

fieldList:
	(names += fieldName ':' types += type ',')* '$' namePattern = IDENT ':' typePattern = type '...' 	# FieldListExpansion
	| (names += fieldName ':' types += type ',')* names += fieldName ':' types += type 					# FieldListNoExpansion
	| 																									# FieldListEmpty;

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