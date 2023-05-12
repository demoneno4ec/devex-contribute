package testcoverage

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestExtractXml(t *testing.T) {
	expect := []Package{
		{
			Path: "src",
			Files: []Coverage{
				{
					File:    "_black_version.py",
					Percent: 100,
				},
			},
		},
		{
			Path: "src.black",
			Files: []Coverage{
				{
					File:           "__init__.py",
					Percent:        83,
					UncoveredLines: []uint32{136, 137, 154, 160, 502, 503, 504, 505, 510, 514, 516, 517, 518, 519, 533, 534, 653, 698, 734, 806, 808, 821, 835, 928, 957, 958, 959, 963, 982, 983, 986, 987, 988, 989, 990, 991, 992, 993, 994, 997, 998, 999, 1000, 1010, 1011, 1012, 1021, 1022, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033, 1034, 1036, 1037, 1038, 1039, 1040, 1041, 1042, 1044, 1311, 1355, 1356, 1357, 1381, 1386, 1413, 1414, 1415, 1416, 1417, 1419, 1420, 1423, 1424, 1425, 1427, 1429, 1430, 1431, 1432, 1433, 1439, 1440, 1442, 1444, 1445, 1449},
				},
			},
		},
	}

	c := make(chan Package, 5)

	err := extractXml(bytes.NewBufferString(testFile), c)
	assert.NoError(t, err)

	var result []Package
	for p := range c {
		result = append(result, p)
	}
	assert.Equal(t, expect, result)
}

func TestDebug(t *testing.T) {
	t.Skip("only for hand testing")

	c := make(chan Package)

	wg, ctx := errgroup.WithContext(context.TODO())

	wg.Go(func() error {
		return ExtractXml(ctx, "/Users/nvrusin/black", c)
	})

	wg.Go(func() error {
		for p := range c {
			t.Log("package", p.Path, len(p.Files))
		}

		return nil
	})

	require.NoError(t, wg.Wait())
}

const testFile = `
<?xml version="1.0" ?>
<coverage version="7.2.5" timestamp="1683648946737" lines-valid="8629" lines-covered="7519" line-rate="0.8714" branches-covered="0" branches-valid="0" branch-rate="0" complexity="0">
	<!-- Generated by coverage.py: https://coverage.readthedocs.io/en/7.2.5 -->
	<!-- Based on https://raw.githubusercontent.com/cobertura/web/master/htdocs/xml/coverage-04.dtd -->
	<sources>
		<source>/Users/nvrusin/black</source>
	</sources>
	<packages>
		<package name="src" line-rate="1" branch-rate="0" complexity="0">
			<classes>
				<class name="_black_version.py" filename="src/_black_version.py" complexity="0" line-rate="1" branch-rate="0">
					<methods/>
					<lines>
						<line number="1" hits="1"/>
					</lines>
				</class>
			</classes>
		</package>
		<package name="src.black" line-rate="0.9215" branch-rate="0" complexity="0">
			<classes>
				<class name="__init__.py" filename="src/black/__init__.py" complexity="0" line-rate="0.831" branch-rate="0">
					<methods/>
					<lines>
						<line number="1" hits="1"/>
						<line number="2" hits="1"/>
						<line number="3" hits="1"/>
						<line number="4" hits="1"/>
						<line number="5" hits="1"/>
						<line number="6" hits="1"/>
						<line number="7" hits="1"/>
						<line number="8" hits="1"/>
						<line number="9" hits="1"/>
						<line number="10" hits="1"/>
						<line number="11" hits="1"/>
						<line number="12" hits="1"/>
						<line number="13" hits="1"/>
						<line number="14" hits="1"/>
						<line number="30" hits="1"/>
						<line number="31" hits="1"/>
						<line number="32" hits="1"/>
						<line number="33" hits="1"/>
						<line number="34" hits="1"/>
						<line number="36" hits="1"/>
						<line number="37" hits="1"/>
						<line number="38" hits="1"/>
						<line number="39" hits="1"/>
						<line number="45" hits="1"/>
						<line number="55" hits="1"/>
						<line number="64" hits="1"/>
						<line number="65" hits="1"/>
						<line number="66" hits="1"/>
						<line number="74" hits="1"/>
						<line number="81" hits="1"/>
						<line number="82" hits="1"/>
						<line number="83" hits="1"/>
						<line number="84" hits="1"/>
						<line number="85" hits="1"/>
						<line number="86" hits="1"/>
						<line number="87" hits="1"/>
						<line number="89" hits="1"/>
						<line number="92" hits="1"/>
						<line number="93" hits="1"/>
						<line number="94" hits="1"/>
						<line number="97" hits="1"/>
						<line number="98" hits="1"/>
						<line number="99" hits="1"/>
						<line number="100" hits="1"/>
						<line number="101" hits="1"/>
						<line number="102" hits="1"/>
						<line number="104" hits="1"/>
						<line number="105" hits="1"/>
						<line number="108" hits="1"/>
						<line number="109" hits="1"/>
						<line number="111" hits="1"/>
						<line number="112" hits="1"/>
						<line number="114" hits="1"/>
						<line number="118" hits="1"/>
						<line number="121" hits="1"/>
						<line number="129" hits="1"/>
						<line number="130" hits="1"/>
						<line number="131" hits="1"/>
						<line number="132" hits="1"/>
						<line number="134" hits="1"/>
						<line number="135" hits="1"/>
						<line number="136" hits="0"/>
						<line number="137" hits="0"/>
						<line number="141" hits="1"/>
						<line number="142" hits="1"/>
						<line number="147" hits="1"/>
						<line number="152" hits="1"/>
						<line number="153" hits="1"/>
						<line number="154" hits="0"/>
						<line number="158" hits="1"/>
						<line number="159" hits="1"/>
						<line number="160" hits="0"/>
						<line number="161" hits="1"/>
						<line number="163" hits="1"/>
						<line number="164" hits="1"/>
						<line number="167" hits="1"/>
						<line number="175" hits="1"/>
						<line number="178" hits="1"/>
						<line number="183" hits="1"/>
						<line number="184" hits="1"/>
						<line number="185" hits="1"/>
						<line number="186" hits="1"/>
						<line number="189" hits="1"/>
						<line number="194" hits="1"/>
						<line number="195" hits="1"/>
						<line number="196" hits="1"/>
						<line number="197" hits="1"/>
						<line number="200" hits="1"/>
						<line number="206" hits="1"/>
						<line number="207" hits="1"/>
						<line number="215" hits="1"/>
						<line number="227" hits="1"/>
						<line number="235" hits="1"/>
						<line number="243" hits="1"/>
						<line number="253" hits="1"/>
						<line number="259" hits="1"/>
						<line number="265" hits="1"/>
						<line number="271" hits="1"/>
						<line number="277" hits="1"/>
						<line number="285" hits="1"/>
						<line number="294" hits="1"/>
						<line number="299" hits="1"/>
						<line number="304" hits="1"/>
						<line number="309" hits="1"/>
						<line number="318" hits="1"/>
						<line number="331" hits="1"/>
						<line number="344" hits="1"/>
						<line number="353" hits="1"/>
						<line number="362" hits="1"/>
						<line number="371" hits="1"/>
						<line number="378" hits="1"/>
						<line number="387" hits="1"/>
						<line number="396" hits="1"/>
						<line number="403" hits="1"/>
						<line number="412" hits="1"/>
						<line number="426" hits="1"/>
						<line number="427" hits="1"/>
						<line number="457" hits="1"/>
						<line number="459" hits="1"/>
						<line number="460" hits="1"/>
						<line number="464" hits="1"/>
						<line number="465" hits="1"/>
						<line number="466" hits="1"/>
						<line number="467" hits="1"/>
						<line number="469" hits="1"/>
						<line number="472" hits="1"/>
						<line number="474" hits="1"/>
						<line number="475" hits="1"/>
						<line number="476" hits="1"/>
						<line number="481" hits="1"/>
						<line number="489" hits="1"/>
						<line number="499" hits="1"/>
						<line number="501" hits="1"/>
						<line number="502" hits="0"/>
						<line number="503" hits="0"/>
						<line number="504" hits="0"/>
						<line number="505" hits="0"/>
						<line number="510" hits="0"/>
						<line number="514" hits="0"/>
						<line number="516" hits="0"/>
						<line number="517" hits="0"/>
						<line number="518" hits="0"/>
						<line number="519" hits="0"/>
						<line number="521" hits="1"/>
						<line number="522" hits="1"/>
						<line number="527" hits="1"/>
						<line number="531" hits="1"/>
						<line number="532" hits="1"/>
						<line number="533" hits="0"/>
						<line number="534" hits="0"/>
						<line number="536" hits="1"/>
						<line number="537" hits="1"/>
						<line number="538" hits="1"/>
						<line number="541" hits="1"/>
						<line number="542" hits="1"/>
						<line number="555" hits="1"/>
						<line number="558" hits="1"/>
						<line number="560" hits="1"/>
						<line number="562" hits="1"/>
						<line number="563" hits="1"/>
						<line number="567" hits="1"/>
						<line number="568" hits="1"/>
						<line number="580" hits="1"/>
						<line number="581" hits="1"/>
						<line number="583" hits="1"/>
						<line number="591" hits="1"/>
						<line number="592" hits="1"/>
						<line number="600" hits="1"/>
						<line number="602" hits="1"/>
						<line number="611" hits="1"/>
						<line number="612" hits="1"/>
						<line number="613" hits="1"/>
						<line number="614" hits="1"/>
						<line number="615" hits="1"/>
						<line number="616" hits="1"/>
						<line number="617" hits="1"/>
						<line number="620" hits="1"/>
						<line number="634" hits="1"/>
						<line number="635" hits="1"/>
						<line number="637" hits="1"/>
						<line number="638" hits="1"/>
						<line number="639" hits="1"/>
						<line number="640" hits="1"/>
						<line number="642" hits="1"/>
						<line number="643" hits="1"/>
						<line number="644" hits="1"/>
						<line number="645" hits="1"/>
						<line number="647" hits="1"/>
						<line number="648" hits="1"/>
						<line number="650" hits="1"/>
						<line number="651" hits="1"/>
						<line number="652" hits="1"/>
						<line number="653" hits="0"/>
						<line number="655" hits="1"/>
						<line number="657" hits="1"/>
						<line number="658" hits="1"/>
						<line number="660" hits="1"/>
						<line number="661" hits="1"/>
						<line number="662" hits="1"/>
						<line number="663" hits="1"/>
						<line number="665" hits="1"/>
						<line number="666" hits="1"/>
						<line number="668" hits="1"/>
						<line number="671" hits="1"/>
						<line number="673" hits="1"/>
						<line number="674" hits="1"/>
						<line number="675" hits="1"/>
						<line number="676" hits="1"/>
						<line number="677" hits="1"/>
						<line number="681" hits="1"/>
						<line number="695" hits="1"/>
						<line number="696" hits="1"/>
						<line number="698" hits="0"/>
						<line number="699" hits="1"/>
						<line number="702" hits="1"/>
						<line number="708" hits="1"/>
						<line number="709" hits="1"/>
						<line number="710" hits="1"/>
						<line number="711" hits="1"/>
						<line number="714" hits="1"/>
						<line number="724" hits="1"/>
						<line number="725" hits="1"/>
						<line number="726" hits="1"/>
						<line number="727" hits="1"/>
						<line number="730" hits="1"/>
						<line number="731" hits="1"/>
						<line number="732" hits="1"/>
						<line number="733" hits="1"/>
						<line number="734" hits="0"/>
						<line number="735" hits="1"/>
						<line number="740" hits="1"/>
						<line number="741" hits="1"/>
						<line number="749" hits="1"/>
						<line number="750" hits="1"/>
						<line number="752" hits="1"/>
						<line number="753" hits="1"/>
						<line number="754" hits="1"/>
						<line number="755" hits="1"/>
						<line number="758" hits="1"/>
						<line number="760" hits="1"/>
						<line number="762" hits="1"/>
						<line number="763" hits="1"/>
						<line number="764" hits="1"/>
						<line number="765" hits="1"/>
						<line number="766" hits="1"/>
						<line number="767" hits="1"/>
						<line number="768" hits="1"/>
						<line number="770" hits="1"/>
						<line number="771" hits="1"/>
						<line number="772" hits="1"/>
						<line number="773" hits="1"/>
						<line number="774" hits="1"/>
						<line number="775" hits="1"/>
						<line number="776" hits="1"/>
						<line number="777" hits="1"/>
						<line number="780" hits="1"/>
						<line number="781" hits="1"/>
						<line number="784" hits="1"/>
						<line number="785" hits="1"/>
						<line number="786" hits="1"/>
						<line number="787" hits="1"/>
						<line number="788" hits="1"/>
						<line number="789" hits="1"/>
						<line number="792" hits="1"/>
						<line number="805" hits="1"/>
						<line number="806" hits="0"/>
						<line number="807" hits="1"/>
						<line number="808" hits="0"/>
						<line number="810" hits="1"/>
						<line number="811" hits="1"/>
						<line number="812" hits="1"/>
						<line number="813" hits="1"/>
						<line number="814" hits="1"/>
						<line number="815" hits="1"/>
						<line number="816" hits="1"/>
						<line number="817" hits="1"/>
						<line number="818" hits="1"/>
						<line number="819" hits="1"/>
						<line number="820" hits="1"/>
						<line number="821" hits="0"/>
						<line number="824" hits="1"/>
						<line number="825" hits="1"/>
						<line number="827" hits="1"/>
						<line number="828" hits="1"/>
						<line number="829" hits="1"/>
						<line number="830" hits="1"/>
						<line number="831" hits="1"/>
						<line number="832" hits="1"/>
						<line number="833" hits="1"/>
						<line number="834" hits="1"/>
						<line number="835" hits="0"/>
						<line number="837" hits="1"/>
						<line number="839" hits="1"/>
						<line number="840" hits="1"/>
						<line number="842" hits="1"/>
						<line number="843" hits="1"/>
						<line number="849" hits="1"/>
						<line number="850" hits="1"/>
						<line number="851" hits="1"/>
						<line number="853" hits="1"/>
						<line number="856" hits="1"/>
						<line number="871" hits="1"/>
						<line number="873" hits="1"/>
						<line number="874" hits="1"/>
						<line number="876" hits="1"/>
						<line number="878" hits="1"/>
						<line number="879" hits="1"/>
						<line number="880" hits="1"/>
						<line number="881" hits="1"/>
						<line number="883" hits="1"/>
						<line number="884" hits="1"/>
						<line number="887" hits="1"/>
						<line number="890" hits="1"/>
						<line number="892" hits="1"/>
						<line number="893" hits="1"/>
						<line number="894" hits="1"/>
						<line number="895" hits="1"/>
						<line number="896" hits="1"/>
						<line number="897" hits="1"/>
						<line number="898" hits="1"/>
						<line number="899" hits="1"/>
						<line number="900" hits="1"/>
						<line number="901" hits="1"/>
						<line number="902" hits="1"/>
						<line number="903" hits="1"/>
						<line number="904" hits="1"/>
						<line number="907" hits="1"/>
						<line number="916" hits="1"/>
						<line number="917" hits="1"/>
						<line number="920" hits="1"/>
						<line number="927" hits="1"/>
						<line number="928" hits="0"/>
						<line number="930" hits="1"/>
						<line number="931" hits="1"/>
						<line number="932" hits="1"/>
						<line number="934" hits="1"/>
						<line number="936" hits="1"/>
						<line number="937" hits="1"/>
						<line number="940" hits="1"/>
						<line number="957" hits="0"/>
						<line number="958" hits="0"/>
						<line number="959" hits="0"/>
						<line number="963" hits="0"/>
						<line number="966" hits="1"/>
						<line number="982" hits="0"/>
						<line number="983" hits="0"/>
						<line number="986" hits="0"/>
						<line number="987" hits="0"/>
						<line number="988" hits="0"/>
						<line number="989" hits="0"/>
						<line number="990" hits="0"/>
						<line number="991" hits="0"/>
						<line number="992" hits="0"/>
						<line number="993" hits="0"/>
						<line number="994" hits="0"/>
						<line number="997" hits="0"/>
						<line number="998" hits="0"/>
						<line number="999" hits="0"/>
						<line number="1000" hits="0"/>
						<line number="1003" hits="1"/>
						<line number="1010" hits="0"/>
						<line number="1011" hits="0"/>
						<line number="1012" hits="0"/>
						<line number="1015" hits="1"/>
						<line number="1021" hits="0"/>
						<line number="1022" hits="0"/>
						<line number="1024" hits="0"/>
						<line number="1025" hits="0"/>
						<line number="1026" hits="0"/>
						<line number="1027" hits="0"/>
						<line number="1028" hits="0"/>
						<line number="1029" hits="0"/>
						<line number="1030" hits="0"/>
						<line number="1031" hits="0"/>
						<line number="1032" hits="0"/>
						<line number="1033" hits="0"/>
						<line number="1034" hits="0"/>
						<line number="1036" hits="0"/>
						<line number="1037" hits="0"/>
						<line number="1038" hits="0"/>
						<line number="1039" hits="0"/>
						<line number="1040" hits="0"/>
						<line number="1041" hits="0"/>
						<line number="1042" hits="0"/>
						<line number="1044" hits="0"/>
						<line number="1047" hits="1"/>
						<line number="1077" hits="1"/>
						<line number="1081" hits="1"/>
						<line number="1082" hits="1"/>
						<line number="1083" hits="1"/>
						<line number="1086" hits="1"/>
						<line number="1087" hits="1"/>
						<line number="1088" hits="1"/>
						<line number="1089" hits="1"/>
						<line number="1090" hits="1"/>
						<line number="1092" hits="1"/>
						<line number="1093" hits="1"/>
						<line number="1095" hits="1"/>
						<line number="1100" hits="1"/>
						<line number="1101" hits="1"/>
						<line number="1102" hits="1"/>
						<line number="1103" hits="1"/>
						<line number="1108" hits="1"/>
						<line number="1109" hits="1"/>
						<line number="1110" hits="1"/>
						<line number="1111" hits="1"/>
						<line number="1112" hits="1"/>
						<line number="1115" hits="1"/>
						<line number="1116" hits="1"/>
						<line number="1117" hits="1"/>
						<line number="1118" hits="1"/>
						<line number="1119" hits="1"/>
						<line number="1120" hits="1"/>
						<line number="1121" hits="1"/>
						<line number="1124" hits="1"/>
						<line number="1125" hits="1"/>
						<line number="1126" hits="1"/>
						<line number="1127" hits="1"/>
						<line number="1128" hits="1"/>
						<line number="1131" hits="1"/>
						<line number="1137" hits="1"/>
						<line number="1138" hits="1"/>
						<line number="1139" hits="1"/>
						<line number="1140" hits="1"/>
						<line number="1142" hits="1"/>
						<line number="1143" hits="1"/>
						<line number="1144" hits="1"/>
						<line number="1145" hits="1"/>
						<line number="1148" hits="1"/>
						<line number="1168" hits="1"/>
						<line number="1169" hits="1"/>
						<line number="1170" hits="1"/>
						<line number="1176" hits="1"/>
						<line number="1177" hits="1"/>
						<line number="1178" hits="1"/>
						<line number="1179" hits="1"/>
						<line number="1180" hits="1"/>
						<line number="1181" hits="1"/>
						<line number="1182" hits="1"/>
						<line number="1183" hits="1"/>
						<line number="1184" hits="1"/>
						<line number="1185" hits="1"/>
						<line number="1187" hits="1"/>
						<line number="1188" hits="1"/>
						<line number="1189" hits="1"/>
						<line number="1191" hits="1"/>
						<line number="1192" hits="1"/>
						<line number="1197" hits="1"/>
						<line number="1199" hits="1"/>
						<line number="1200" hits="1"/>
						<line number="1202" hits="1"/>
						<line number="1203" hits="1"/>
						<line number="1206" hits="1"/>
						<line number="1208" hits="1"/>
						<line number="1213" hits="1"/>
						<line number="1214" hits="1"/>
						<line number="1216" hits="1"/>
						<line number="1218" hits="1"/>
						<line number="1219" hits="1"/>
						<line number="1220" hits="1"/>
						<line number="1222" hits="1"/>
						<line number="1223" hits="1"/>
						<line number="1224" hits="1"/>
						<line number="1225" hits="1"/>
						<line number="1227" hits="1"/>
						<line number="1233" hits="1"/>
						<line number="1235" hits="1"/>
						<line number="1240" hits="1"/>
						<line number="1242" hits="1"/>
						<line number="1247" hits="1"/>
						<line number="1248" hits="1"/>
						<line number="1254" hits="1"/>
						<line number="1256" hits="1"/>
						<line number="1257" hits="1"/>
						<line number="1259" hits="1"/>
						<line number="1264" hits="1"/>
						<line number="1266" hits="1"/>
						<line number="1269" hits="1"/>
						<line number="1271" hits="1"/>
						<line number="1276" hits="1"/>
						<line number="1278" hits="1"/>
						<line number="1281" hits="1"/>
						<line number="1285" hits="1"/>
						<line number="1286" hits="1"/>
						<line number="1291" hits="1"/>
						<line number="1293" hits="1"/>
						<line number="1295" hits="1"/>
						<line number="1296" hits="1"/>
						<line number="1297" hits="1"/>
						<line number="1298" hits="1"/>
						<line number="1299" hits="1"/>
						<line number="1301" hits="1"/>
						<line number="1302" hits="1"/>
						<line number="1303" hits="1"/>
						<line number="1304" hits="1"/>
						<line number="1305" hits="1"/>
						<line number="1307" hits="1"/>
						<line number="1308" hits="1"/>
						<line number="1311" hits="0"/>
						<line number="1313" hits="1"/>
						<line number="1314" hits="1"/>
						<line number="1315" hits="1"/>
						<line number="1317" hits="1"/>
						<line number="1318" hits="1"/>
						<line number="1320" hits="1"/>
						<line number="1325" hits="1"/>
						<line number="1327" hits="1"/>
						<line number="1329" hits="1"/>
						<line number="1330" hits="1"/>
						<line number="1331" hits="1"/>
						<line number="1332" hits="1"/>
						<line number="1334" hits="1"/>
						<line number="1336" hits="1"/>
						<line number="1338" hits="1"/>
						<line number="1341" hits="1"/>
						<line number="1343" hits="1"/>
						<line number="1344" hits="1"/>
						<line number="1345" hits="1"/>
						<line number="1346" hits="1"/>
						<line number="1353" hits="1"/>
						<line number="1354" hits="1"/>
						<line number="1355" hits="0"/>
						<line number="1356" hits="0"/>
						<line number="1357" hits="0"/>
						<line number="1363" hits="1"/>
						<line number="1364" hits="1"/>
						<line number="1365" hits="1"/>
						<line number="1366" hits="1"/>
						<line number="1367" hits="1"/>
						<line number="1374" hits="1"/>
						<line number="1379" hits="1"/>
						<line number="1380" hits="1"/>
						<line number="1381" hits="0"/>
						<line number="1386" hits="0"/>
						<line number="1393" hits="1"/>
						<line number="1394" hits="1"/>
						<line number="1399" hits="1"/>
						<line number="1402" hits="1"/>
						<line number="1413" hits="0"/>
						<line number="1414" hits="0"/>
						<line number="1415" hits="0"/>
						<line number="1416" hits="0"/>
						<line number="1417" hits="0"/>
						<line number="1419" hits="0"/>
						<line number="1420" hits="0"/>
						<line number="1423" hits="0"/>
						<line number="1424" hits="0"/>
						<line number="1425" hits="0"/>
						<line number="1427" hits="0"/>
						<line number="1429" hits="0"/>
						<line number="1430" hits="0"/>
						<line number="1431" hits="0"/>
						<line number="1432" hits="0"/>
						<line number="1433" hits="0"/>
						<line number="1436" hits="1"/>
						<line number="1439" hits="0"/>
						<line number="1440" hits="0"/>
						<line number="1442" hits="0"/>
						<line number="1444" hits="0"/>
						<line number="1445" hits="0"/>
						<line number="1448" hits="1"/>
						<line number="1449" hits="0"/>
					</lines>
				</class>
			</classes>
		</package>
	</packages>
</coverage>
`
