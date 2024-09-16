// Generated from /home/htao/sdp-lsp/ast/parser/Sdp.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Sdp extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, NEW_LINE=2, LINE_COMMENT=3, BLOCK_COMMENT=4, MODULE=5, HASH=6, INCLUDE=7, 
		STRUCT=8, ENUM=9, OPTIONAL=10, REQUIRED=11, VECTOR=12, MAP=13, ULONG=14, 
		LONG=15, FLOAT=16, BYTE=17, UINT=18, INT=19, DOUBLE=20, STRING=21, VOID=22, 
		BOOL=23, TRUE=24, FALSE=25, INTERFACE=26, IN=27, OUT=28, LEFT_BRACE=29, 
		RIGHT_BRACE=30, LEFT_BRACKET=31, RIGHT_BRACKET=32, LEFT_PAREN=33, RIGHT_PAREN=34, 
		SEMICOLON=35, COLON=36, COMMA=37, EQUAL=38, LEFT_ANGLE_BRACKET=39, RIGHT_ANGLE_BRACKET=40, 
		STRING_LITERAL=41, NUMBER=42, IDENTIFIER=43, DOUBLE_COLON=44;
	public static final int
		RULE_sdp = 0, RULE_includeStatement = 1, RULE_moduleStatement = 2, RULE_structStatement = 3, 
		RULE_fieldStatement = 4, RULE_serviceStatement = 5, RULE_serviceMethodStatement = 6, 
		RULE_methodParamStatement = 7, RULE_enumStatement = 8, RULE_enumField = 9, 
		RULE_enumLastField = 10, RULE_type = 11, RULE_vectorType = 12, RULE_mapType = 13, 
		RULE_moduleType = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"sdp", "includeStatement", "moduleStatement", "structStatement", "fieldStatement", 
			"serviceStatement", "serviceMethodStatement", "methodParamStatement", 
			"enumStatement", "enumField", "enumLastField", "type", "vectorType", 
			"mapType", "moduleType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, "'module'", "'#'", "'include'", "'struct'", 
			"'enum'", "'optional'", "'required'", "'vector'", "'map'", "'unsigned long'", 
			"'long'", "'float'", "'byte'", "'unsigned int'", "'int'", "'double'", 
			"'string'", "'void'", "'bool'", "'true'", "'false'", "'interface'", "'in'", 
			"'out'", "'{'", "'}'", "'['", "']'", "'('", "')'", "';'", "':'", "','", 
			"'='", "'<'", "'>'", null, null, null, "'::'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "NEW_LINE", "LINE_COMMENT", "BLOCK_COMMENT", "MODULE", "HASH", 
			"INCLUDE", "STRUCT", "ENUM", "OPTIONAL", "REQUIRED", "VECTOR", "MAP", 
			"ULONG", "LONG", "FLOAT", "BYTE", "UINT", "INT", "DOUBLE", "STRING", 
			"VOID", "BOOL", "TRUE", "FALSE", "INTERFACE", "IN", "OUT", "LEFT_BRACE", 
			"RIGHT_BRACE", "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_PAREN", "RIGHT_PAREN", 
			"SEMICOLON", "COLON", "COMMA", "EQUAL", "LEFT_ANGLE_BRACKET", "RIGHT_ANGLE_BRACKET", 
			"STRING_LITERAL", "NUMBER", "IDENTIFIER", "DOUBLE_COLON"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Sdp.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Sdp(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SdpContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(Sdp.EOF, 0); }
		public List<IncludeStatementContext> includeStatement() {
			return getRuleContexts(IncludeStatementContext.class);
		}
		public IncludeStatementContext includeStatement(int i) {
			return getRuleContext(IncludeStatementContext.class,i);
		}
		public List<ModuleStatementContext> moduleStatement() {
			return getRuleContexts(ModuleStatementContext.class);
		}
		public ModuleStatementContext moduleStatement(int i) {
			return getRuleContext(ModuleStatementContext.class,i);
		}
		public SdpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sdp; }
	}

	public final SdpContext sdp() throws RecognitionException {
		SdpContext _localctx = new SdpContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_sdp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==HASH) {
				{
				{
				setState(30);
				includeStatement();
				}
				}
				setState(35);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(39);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MODULE) {
				{
				{
				setState(36);
				moduleStatement();
				}
				}
				setState(41);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(42);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IncludeStatementContext extends ParserRuleContext {
		public TerminalNode HASH() { return getToken(Sdp.HASH, 0); }
		public TerminalNode INCLUDE() { return getToken(Sdp.INCLUDE, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(Sdp.STRING_LITERAL, 0); }
		public IncludeStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_includeStatement; }
	}

	public final IncludeStatementContext includeStatement() throws RecognitionException {
		IncludeStatementContext _localctx = new IncludeStatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_includeStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(HASH);
			setState(45);
			match(INCLUDE);
			setState(46);
			match(STRING_LITERAL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleStatementContext extends ParserRuleContext {
		public TerminalNode MODULE() { return getToken(Sdp.MODULE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode LEFT_BRACE() { return getToken(Sdp.LEFT_BRACE, 0); }
		public TerminalNode RIGHT_BRACE() { return getToken(Sdp.RIGHT_BRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public List<StructStatementContext> structStatement() {
			return getRuleContexts(StructStatementContext.class);
		}
		public StructStatementContext structStatement(int i) {
			return getRuleContext(StructStatementContext.class,i);
		}
		public List<ServiceStatementContext> serviceStatement() {
			return getRuleContexts(ServiceStatementContext.class);
		}
		public ServiceStatementContext serviceStatement(int i) {
			return getRuleContext(ServiceStatementContext.class,i);
		}
		public List<EnumStatementContext> enumStatement() {
			return getRuleContexts(EnumStatementContext.class);
		}
		public EnumStatementContext enumStatement(int i) {
			return getRuleContext(EnumStatementContext.class,i);
		}
		public ModuleStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleStatement; }
	}

	public final ModuleStatementContext moduleStatement() throws RecognitionException {
		ModuleStatementContext _localctx = new ModuleStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_moduleStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			match(MODULE);
			setState(49);
			match(IDENTIFIER);
			setState(50);
			match(LEFT_BRACE);
			setState(56);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 67109632L) != 0)) {
				{
				setState(54);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case STRUCT:
					{
					setState(51);
					structStatement();
					}
					break;
				case INTERFACE:
					{
					setState(52);
					serviceStatement();
					}
					break;
				case ENUM:
					{
					setState(53);
					enumStatement();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(59);
			match(RIGHT_BRACE);
			setState(60);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructStatementContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(Sdp.STRUCT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode LEFT_BRACE() { return getToken(Sdp.LEFT_BRACE, 0); }
		public TerminalNode RIGHT_BRACE() { return getToken(Sdp.RIGHT_BRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public List<FieldStatementContext> fieldStatement() {
			return getRuleContexts(FieldStatementContext.class);
		}
		public FieldStatementContext fieldStatement(int i) {
			return getRuleContext(FieldStatementContext.class,i);
		}
		public StructStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structStatement; }
	}

	public final StructStatementContext structStatement() throws RecognitionException {
		StructStatementContext _localctx = new StructStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_structStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(STRUCT);
			setState(63);
			match(IDENTIFIER);
			setState(64);
			match(LEFT_BRACE);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==NUMBER) {
				{
				{
				setState(65);
				fieldStatement();
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(71);
			match(RIGHT_BRACE);
			setState(72);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldStatementContext extends ParserRuleContext {
		public List<TerminalNode> NUMBER() { return getTokens(Sdp.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(Sdp.NUMBER, i);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public TerminalNode OPTIONAL() { return getToken(Sdp.OPTIONAL, 0); }
		public TerminalNode REQUIRED() { return getToken(Sdp.REQUIRED, 0); }
		public TerminalNode EQUAL() { return getToken(Sdp.EQUAL, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(Sdp.STRING_LITERAL, 0); }
		public TerminalNode TRUE() { return getToken(Sdp.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Sdp.FALSE, 0); }
		public FieldStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldStatement; }
	}

	public final FieldStatementContext fieldStatement() throws RecognitionException {
		FieldStatementContext _localctx = new FieldStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_fieldStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			match(NUMBER);
			setState(75);
			_la = _input.LA(1);
			if ( !(_la==OPTIONAL || _la==REQUIRED) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(76);
			type();
			setState(77);
			match(IDENTIFIER);
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(78);
				match(EQUAL);
				setState(79);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 6597120098304L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(82);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ServiceStatementContext extends ParserRuleContext {
		public TerminalNode INTERFACE() { return getToken(Sdp.INTERFACE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode LEFT_BRACE() { return getToken(Sdp.LEFT_BRACE, 0); }
		public TerminalNode RIGHT_BRACE() { return getToken(Sdp.RIGHT_BRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public List<ServiceMethodStatementContext> serviceMethodStatement() {
			return getRuleContexts(ServiceMethodStatementContext.class);
		}
		public ServiceMethodStatementContext serviceMethodStatement(int i) {
			return getRuleContext(ServiceMethodStatementContext.class,i);
		}
		public ServiceStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_serviceStatement; }
	}

	public final ServiceStatementContext serviceStatement() throws RecognitionException {
		ServiceStatementContext _localctx = new ServiceStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_serviceStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			match(INTERFACE);
			setState(85);
			match(IDENTIFIER);
			setState(86);
			match(LEFT_BRACE);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8796109795328L) != 0)) {
				{
				{
				setState(87);
				serviceMethodStatement();
				}
				}
				setState(92);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(93);
			match(RIGHT_BRACE);
			setState(94);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ServiceMethodStatementContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode LEFT_PAREN() { return getToken(Sdp.LEFT_PAREN, 0); }
		public TerminalNode RIGHT_PAREN() { return getToken(Sdp.RIGHT_PAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public MethodParamStatementContext methodParamStatement() {
			return getRuleContext(MethodParamStatementContext.class,0);
		}
		public ServiceMethodStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_serviceMethodStatement; }
	}

	public final ServiceMethodStatementContext serviceMethodStatement() throws RecognitionException {
		ServiceMethodStatementContext _localctx = new ServiceMethodStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_serviceMethodStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			type();
			setState(97);
			match(IDENTIFIER);
			setState(98);
			match(LEFT_PAREN);
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8796512448512L) != 0)) {
				{
				setState(99);
				methodParamStatement();
				}
			}

			setState(102);
			match(RIGHT_PAREN);
			setState(103);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodParamStatementContext extends ParserRuleContext {
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public List<TerminalNode> IDENTIFIER() { return getTokens(Sdp.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(Sdp.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(Sdp.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(Sdp.COMMA, i);
		}
		public List<TerminalNode> IN() { return getTokens(Sdp.IN); }
		public TerminalNode IN(int i) {
			return getToken(Sdp.IN, i);
		}
		public List<TerminalNode> OUT() { return getTokens(Sdp.OUT); }
		public TerminalNode OUT(int i) {
			return getToken(Sdp.OUT, i);
		}
		public MethodParamStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodParamStatement; }
	}

	public final MethodParamStatementContext methodParamStatement() throws RecognitionException {
		MethodParamStatementContext _localctx = new MethodParamStatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_methodParamStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IN || _la==OUT) {
				{
				setState(105);
				_la = _input.LA(1);
				if ( !(_la==IN || _la==OUT) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(108);
			type();
			setState(109);
			match(IDENTIFIER);
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(110);
				match(COMMA);
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IN || _la==OUT) {
					{
					setState(111);
					_la = _input.LA(1);
					if ( !(_la==IN || _la==OUT) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

				setState(114);
				type();
				setState(115);
				match(IDENTIFIER);
				}
				}
				setState(121);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumStatementContext extends ParserRuleContext {
		public TerminalNode ENUM() { return getToken(Sdp.ENUM, 0); }
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode LEFT_BRACE() { return getToken(Sdp.LEFT_BRACE, 0); }
		public TerminalNode RIGHT_BRACE() { return getToken(Sdp.RIGHT_BRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(Sdp.SEMICOLON, 0); }
		public List<EnumFieldContext> enumField() {
			return getRuleContexts(EnumFieldContext.class);
		}
		public EnumFieldContext enumField(int i) {
			return getRuleContext(EnumFieldContext.class,i);
		}
		public EnumStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumStatement; }
	}

	public final EnumStatementContext enumStatement() throws RecognitionException {
		EnumStatementContext _localctx = new EnumStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_enumStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(ENUM);
			setState(123);
			match(IDENTIFIER);
			setState(124);
			match(LEFT_BRACE);
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(125);
				enumField();
				}
				}
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(131);
			match(RIGHT_BRACE);
			setState(132);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumFieldContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode COMMA() { return getToken(Sdp.COMMA, 0); }
		public TerminalNode EQUAL() { return getToken(Sdp.EQUAL, 0); }
		public TerminalNode NUMBER() { return getToken(Sdp.NUMBER, 0); }
		public EnumFieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumField; }
	}

	public final EnumFieldContext enumField() throws RecognitionException {
		EnumFieldContext _localctx = new EnumFieldContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_enumField);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(IDENTIFIER);
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(135);
				match(EQUAL);
				setState(136);
				match(NUMBER);
				}
			}

			setState(139);
			match(COMMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumLastFieldContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public TerminalNode COMMA() { return getToken(Sdp.COMMA, 0); }
		public EnumLastFieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumLastField; }
	}

	public final EnumLastFieldContext enumLastField() throws RecognitionException {
		EnumLastFieldContext _localctx = new EnumLastFieldContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_enumLastField);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(IDENTIFIER);
			setState(142);
			match(COMMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode UINT() { return getToken(Sdp.UINT, 0); }
		public TerminalNode ULONG() { return getToken(Sdp.ULONG, 0); }
		public TerminalNode LONG() { return getToken(Sdp.LONG, 0); }
		public TerminalNode FLOAT() { return getToken(Sdp.FLOAT, 0); }
		public TerminalNode BYTE() { return getToken(Sdp.BYTE, 0); }
		public TerminalNode INT() { return getToken(Sdp.INT, 0); }
		public TerminalNode DOUBLE() { return getToken(Sdp.DOUBLE, 0); }
		public TerminalNode BOOL() { return getToken(Sdp.BOOL, 0); }
		public TerminalNode STRING() { return getToken(Sdp.STRING, 0); }
		public TerminalNode VOID() { return getToken(Sdp.VOID, 0); }
		public TerminalNode IDENTIFIER() { return getToken(Sdp.IDENTIFIER, 0); }
		public ModuleTypeContext moduleType() {
			return getRuleContext(ModuleTypeContext.class,0);
		}
		public VectorTypeContext vectorType() {
			return getRuleContext(VectorTypeContext.class,0);
		}
		public MapTypeContext mapType() {
			return getRuleContext(MapTypeContext.class,0);
		}
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_type);
		try {
			setState(158);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(144);
				match(UINT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(145);
				match(ULONG);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(146);
				match(LONG);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(147);
				match(FLOAT);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(148);
				match(BYTE);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(149);
				match(INT);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(150);
				match(DOUBLE);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(151);
				match(BOOL);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(152);
				match(STRING);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(153);
				match(VOID);
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(154);
				match(IDENTIFIER);
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(155);
				moduleType();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(156);
				vectorType();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(157);
				mapType();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorTypeContext extends ParserRuleContext {
		public TerminalNode VECTOR() { return getToken(Sdp.VECTOR, 0); }
		public TerminalNode LEFT_ANGLE_BRACKET() { return getToken(Sdp.LEFT_ANGLE_BRACKET, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode RIGHT_ANGLE_BRACKET() { return getToken(Sdp.RIGHT_ANGLE_BRACKET, 0); }
		public VectorTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorType; }
	}

	public final VectorTypeContext vectorType() throws RecognitionException {
		VectorTypeContext _localctx = new VectorTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_vectorType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(VECTOR);
			setState(161);
			match(LEFT_ANGLE_BRACKET);
			setState(162);
			type();
			setState(163);
			match(RIGHT_ANGLE_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MapTypeContext extends ParserRuleContext {
		public TerminalNode MAP() { return getToken(Sdp.MAP, 0); }
		public TerminalNode LEFT_ANGLE_BRACKET() { return getToken(Sdp.LEFT_ANGLE_BRACKET, 0); }
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(Sdp.COMMA, 0); }
		public TerminalNode RIGHT_ANGLE_BRACKET() { return getToken(Sdp.RIGHT_ANGLE_BRACKET, 0); }
		public MapTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mapType; }
	}

	public final MapTypeContext mapType() throws RecognitionException {
		MapTypeContext _localctx = new MapTypeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_mapType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			match(MAP);
			setState(166);
			match(LEFT_ANGLE_BRACKET);
			setState(167);
			type();
			setState(168);
			match(COMMA);
			setState(169);
			type();
			setState(170);
			match(RIGHT_ANGLE_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleTypeContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(Sdp.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(Sdp.IDENTIFIER, i);
		}
		public TerminalNode DOUBLE_COLON() { return getToken(Sdp.DOUBLE_COLON, 0); }
		public ModuleTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleType; }
	}

	public final ModuleTypeContext moduleType() throws RecognitionException {
		ModuleTypeContext _localctx = new ModuleTypeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_moduleType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			match(IDENTIFIER);
			setState(173);
			match(DOUBLE_COLON);
			setState(174);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001,\u00b1\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001\u0000\u0005\u0000"+
		" \b\u0000\n\u0000\f\u0000#\t\u0000\u0001\u0000\u0005\u0000&\b\u0000\n"+
		"\u0000\f\u0000)\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0005\u00027\b\u0002\n\u0002\f\u0002:\t\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0005\u0003C\b\u0003\n\u0003\f\u0003F\t\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004Q\b\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005Y\b\u0005"+
		"\n\u0005\f\u0005\\\t\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006e\b\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0003\u0007k\b\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007q\b\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0005\u0007v\b\u0007\n\u0007\f\u0007y\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0005\b\u007f\b\b\n\b\f\b\u0082\t\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0003\t\u008a\b\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u009f\b\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0000\u0000\u000f\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u0000\u0003\u0001\u0000\n\u000b\u0002"+
		"\u0000\u0018\u0019)*\u0001\u0000\u001b\u001c\u00bc\u0000!\u0001\u0000"+
		"\u0000\u0000\u0002,\u0001\u0000\u0000\u0000\u00040\u0001\u0000\u0000\u0000"+
		"\u0006>\u0001\u0000\u0000\u0000\bJ\u0001\u0000\u0000\u0000\nT\u0001\u0000"+
		"\u0000\u0000\f`\u0001\u0000\u0000\u0000\u000ej\u0001\u0000\u0000\u0000"+
		"\u0010z\u0001\u0000\u0000\u0000\u0012\u0086\u0001\u0000\u0000\u0000\u0014"+
		"\u008d\u0001\u0000\u0000\u0000\u0016\u009e\u0001\u0000\u0000\u0000\u0018"+
		"\u00a0\u0001\u0000\u0000\u0000\u001a\u00a5\u0001\u0000\u0000\u0000\u001c"+
		"\u00ac\u0001\u0000\u0000\u0000\u001e \u0003\u0002\u0001\u0000\u001f\u001e"+
		"\u0001\u0000\u0000\u0000 #\u0001\u0000\u0000\u0000!\u001f\u0001\u0000"+
		"\u0000\u0000!\"\u0001\u0000\u0000\u0000\"\'\u0001\u0000\u0000\u0000#!"+
		"\u0001\u0000\u0000\u0000$&\u0003\u0004\u0002\u0000%$\u0001\u0000\u0000"+
		"\u0000&)\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000\u0000\'(\u0001\u0000"+
		"\u0000\u0000(*\u0001\u0000\u0000\u0000)\'\u0001\u0000\u0000\u0000*+\u0005"+
		"\u0000\u0000\u0001+\u0001\u0001\u0000\u0000\u0000,-\u0005\u0006\u0000"+
		"\u0000-.\u0005\u0007\u0000\u0000./\u0005)\u0000\u0000/\u0003\u0001\u0000"+
		"\u0000\u000001\u0005\u0005\u0000\u000012\u0005+\u0000\u000028\u0005\u001d"+
		"\u0000\u000037\u0003\u0006\u0003\u000047\u0003\n\u0005\u000057\u0003\u0010"+
		"\b\u000063\u0001\u0000\u0000\u000064\u0001\u0000\u0000\u000065\u0001\u0000"+
		"\u0000\u00007:\u0001\u0000\u0000\u000086\u0001\u0000\u0000\u000089\u0001"+
		"\u0000\u0000\u00009;\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000\u0000"+
		";<\u0005\u001e\u0000\u0000<=\u0005#\u0000\u0000=\u0005\u0001\u0000\u0000"+
		"\u0000>?\u0005\b\u0000\u0000?@\u0005+\u0000\u0000@D\u0005\u001d\u0000"+
		"\u0000AC\u0003\b\u0004\u0000BA\u0001\u0000\u0000\u0000CF\u0001\u0000\u0000"+
		"\u0000DB\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000EG\u0001\u0000"+
		"\u0000\u0000FD\u0001\u0000\u0000\u0000GH\u0005\u001e\u0000\u0000HI\u0005"+
		"#\u0000\u0000I\u0007\u0001\u0000\u0000\u0000JK\u0005*\u0000\u0000KL\u0007"+
		"\u0000\u0000\u0000LM\u0003\u0016\u000b\u0000MP\u0005+\u0000\u0000NO\u0005"+
		"&\u0000\u0000OQ\u0007\u0001\u0000\u0000PN\u0001\u0000\u0000\u0000PQ\u0001"+
		"\u0000\u0000\u0000QR\u0001\u0000\u0000\u0000RS\u0005#\u0000\u0000S\t\u0001"+
		"\u0000\u0000\u0000TU\u0005\u001a\u0000\u0000UV\u0005+\u0000\u0000VZ\u0005"+
		"\u001d\u0000\u0000WY\u0003\f\u0006\u0000XW\u0001\u0000\u0000\u0000Y\\"+
		"\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000Z[\u0001\u0000\u0000"+
		"\u0000[]\u0001\u0000\u0000\u0000\\Z\u0001\u0000\u0000\u0000]^\u0005\u001e"+
		"\u0000\u0000^_\u0005#\u0000\u0000_\u000b\u0001\u0000\u0000\u0000`a\u0003"+
		"\u0016\u000b\u0000ab\u0005+\u0000\u0000bd\u0005!\u0000\u0000ce\u0003\u000e"+
		"\u0007\u0000dc\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000ef\u0001"+
		"\u0000\u0000\u0000fg\u0005\"\u0000\u0000gh\u0005#\u0000\u0000h\r\u0001"+
		"\u0000\u0000\u0000ik\u0007\u0002\u0000\u0000ji\u0001\u0000\u0000\u0000"+
		"jk\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lm\u0003\u0016\u000b"+
		"\u0000mw\u0005+\u0000\u0000np\u0005%\u0000\u0000oq\u0007\u0002\u0000\u0000"+
		"po\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000"+
		"\u0000rs\u0003\u0016\u000b\u0000st\u0005+\u0000\u0000tv\u0001\u0000\u0000"+
		"\u0000un\u0001\u0000\u0000\u0000vy\u0001\u0000\u0000\u0000wu\u0001\u0000"+
		"\u0000\u0000wx\u0001\u0000\u0000\u0000x\u000f\u0001\u0000\u0000\u0000"+
		"yw\u0001\u0000\u0000\u0000z{\u0005\t\u0000\u0000{|\u0005+\u0000\u0000"+
		"|\u0080\u0005\u001d\u0000\u0000}\u007f\u0003\u0012\t\u0000~}\u0001\u0000"+
		"\u0000\u0000\u007f\u0082\u0001\u0000\u0000\u0000\u0080~\u0001\u0000\u0000"+
		"\u0000\u0080\u0081\u0001\u0000\u0000\u0000\u0081\u0083\u0001\u0000\u0000"+
		"\u0000\u0082\u0080\u0001\u0000\u0000\u0000\u0083\u0084\u0005\u001e\u0000"+
		"\u0000\u0084\u0085\u0005#\u0000\u0000\u0085\u0011\u0001\u0000\u0000\u0000"+
		"\u0086\u0089\u0005+\u0000\u0000\u0087\u0088\u0005&\u0000\u0000\u0088\u008a"+
		"\u0005*\u0000\u0000\u0089\u0087\u0001\u0000\u0000\u0000\u0089\u008a\u0001"+
		"\u0000\u0000\u0000\u008a\u008b\u0001\u0000\u0000\u0000\u008b\u008c\u0005"+
		"%\u0000\u0000\u008c\u0013\u0001\u0000\u0000\u0000\u008d\u008e\u0005+\u0000"+
		"\u0000\u008e\u008f\u0005%\u0000\u0000\u008f\u0015\u0001\u0000\u0000\u0000"+
		"\u0090\u009f\u0005\u0012\u0000\u0000\u0091\u009f\u0005\u000e\u0000\u0000"+
		"\u0092\u009f\u0005\u000f\u0000\u0000\u0093\u009f\u0005\u0010\u0000\u0000"+
		"\u0094\u009f\u0005\u0011\u0000\u0000\u0095\u009f\u0005\u0013\u0000\u0000"+
		"\u0096\u009f\u0005\u0014\u0000\u0000\u0097\u009f\u0005\u0017\u0000\u0000"+
		"\u0098\u009f\u0005\u0015\u0000\u0000\u0099\u009f\u0005\u0016\u0000\u0000"+
		"\u009a\u009f\u0005+\u0000\u0000\u009b\u009f\u0003\u001c\u000e\u0000\u009c"+
		"\u009f\u0003\u0018\f\u0000\u009d\u009f\u0003\u001a\r\u0000\u009e\u0090"+
		"\u0001\u0000\u0000\u0000\u009e\u0091\u0001\u0000\u0000\u0000\u009e\u0092"+
		"\u0001\u0000\u0000\u0000\u009e\u0093\u0001\u0000\u0000\u0000\u009e\u0094"+
		"\u0001\u0000\u0000\u0000\u009e\u0095\u0001\u0000\u0000\u0000\u009e\u0096"+
		"\u0001\u0000\u0000\u0000\u009e\u0097\u0001\u0000\u0000\u0000\u009e\u0098"+
		"\u0001\u0000\u0000\u0000\u009e\u0099\u0001\u0000\u0000\u0000\u009e\u009a"+
		"\u0001\u0000\u0000\u0000\u009e\u009b\u0001\u0000\u0000\u0000\u009e\u009c"+
		"\u0001\u0000\u0000\u0000\u009e\u009d\u0001\u0000\u0000\u0000\u009f\u0017"+
		"\u0001\u0000\u0000\u0000\u00a0\u00a1\u0005\f\u0000\u0000\u00a1\u00a2\u0005"+
		"\'\u0000\u0000\u00a2\u00a3\u0003\u0016\u000b\u0000\u00a3\u00a4\u0005("+
		"\u0000\u0000\u00a4\u0019\u0001\u0000\u0000\u0000\u00a5\u00a6\u0005\r\u0000"+
		"\u0000\u00a6\u00a7\u0005\'\u0000\u0000\u00a7\u00a8\u0003\u0016\u000b\u0000"+
		"\u00a8\u00a9\u0005%\u0000\u0000\u00a9\u00aa\u0003\u0016\u000b\u0000\u00aa"+
		"\u00ab\u0005(\u0000\u0000\u00ab\u001b\u0001\u0000\u0000\u0000\u00ac\u00ad"+
		"\u0005+\u0000\u0000\u00ad\u00ae\u0005,\u0000\u0000\u00ae\u00af\u0005+"+
		"\u0000\u0000\u00af\u001d\u0001\u0000\u0000\u0000\u000e!\'68DPZdjpw\u0080"+
		"\u0089\u009e";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}