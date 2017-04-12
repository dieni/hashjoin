package mylib

import (
	"fmt"
	"os"
	"os/exec"
)

type Course struct {
	Id   string
	Name string
}

type Student struct {
	Id       int
	CourseId string
}

var courses []Course
var students []Student

const minCoursesPerStudent = 3
const maxCoursesPerStudent = 5

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

//Generates a table of students of a given size
func CreateStudentTable(size int) []Student {

	students = make([]Student, size)

	var i int = 0
	for studentId := 1; i < size; studentId++ {
		var qCourses = random(minCoursesPerStudent, maxCoursesPerStudent) //select random quantity of courses
		ranCourses := make([]string, qCourses)                            //create a slice of type string

		var k int = 0
		for k < qCourses { //add courses as long the slice is not full

			ranCourseId := courses[random(0, len(courses))].Id //Select a random course

			if !containsCourse(ranCourses, ranCourseId) { //see if course already in slice and if not, add it
				ranCourses[k] = ranCourseId
				k++
			}
		}

		for _, c := range ranCourses {
			if i < size {
				students[i] = Student{Id: studentId, CourseId: c}
				i++
			} else {
				return students
			}
		}
		fmt.Println("Building student table:", i*100/size, " %") //Progress
	}

	return students
}

func CreateCourseTable() []Course {
	courses = make([]Course, 375)
	courses[0] = Course{Id: "051010", Name: "VU Programmierung 1"}
courses[1] = Course{Id: "051011", Name: "VO Technische Grundlagen der Informatik"}
courses[2] = Course{Id: "051012", Name: "PUE Repetitorium Technische Grundlagen der Informatik"}
courses[3] = Course{Id: "051110", Name: "VO Mathematische Grundlagen der Informatik 1"}
courses[4] = Course{Id: "051111", Name: "PUE Repetitorium Mathematische Grundlagen der Informatik 1"}
courses[5] = Course{Id: "051013", Name: "VO Theoretische Informatik"}
courses[6] = Course{Id: "051014", Name: "PUE Repetitorium Theoretische Informatik"}
courses[7] = Course{Id: "051020", Name: "VU Programmierung 2"}
courses[8] = Course{Id: "051024", Name: "VU Algorithmen und Datenstrukturen 1"}
courses[9] = Course{Id: "051034", Name: "VO Netzwerktechnologien"}
courses[10] = Course{Id: "051035", Name: "UE Netzwerktechnologien"}
courses[11] = Course{Id: "051050", Name: "VUSoftware Engineering 2"}
courses[12] = Course{Id: "051039", Name: "VU Projektmanagement"}
courses[13] = Course{Id: "051041", Name: "VU Mensch-Computer-Interaktion"}
courses[14] = Course{Id: "051120", Name: "VO Mathematische Grundlagen der Informatik 2"}
courses[15] = Course{Id: "051121", Name: "UE Mathematische Grundlagen der Informatik 2"}
courses[16] = Course{Id: "051130", Name: "VO Einführende Statistik"}
courses[17] = Course{Id: "051131", Name: "UE Einführende Statistik"}
courses[18] = Course{Id: "040686", Name: "KU Repetitorium: Grundzüge der ABWL"}
courses[19] = Course{Id: "051210", Name: "VO Grundlagen der Wirtschaftsinformatik"}
courses[20] = Course{Id: "051023", Name: "VU Modellierung"}
courses[21] = Course{Id: "051061", Name: "VU Informationssicherheit"}
courses[22] = Course{Id: "051261", Name: "VU Enterprise Architecture: Design, Integration, Implementierung"}
courses[23] = Course{Id: "051240", Name: "VU Unternehmensinformationssysteme"}
courses[24] = Course{Id: "052400", Name: "VU Information Management & Systems Engineering"}
courses[25] = Course{Id: "051019", Name: "VO Informatik und Recht"}
courses[26] = Course{Id: "051029", Name: "VU Informatik und Gesellschaft"}
courses[27] = Course{Id: "040078", Name: "VO Grundzüge der Volkswirtschaftslehre"}
courses[28] = Course{Id: "040136", Name: "VO ABWL Produktion und Logistik I"}
courses[29] = Course{Id: "040137", Name: "UK ABWL Produktion und Logistik II"}
courses[30] = Course{Id: "040002", Name: "VO ABWL Finanzwirtschaft I"}
courses[31] = Course{Id: "040003", Name: "UK ABWL Finanzwirtschaft II"}
courses[32] = Course{Id: "040373", Name: "VO Grundzüge des Rechts"}
courses[33] = Course{Id: "040712", Name: "UK Unternehmensrecht"}
courses[34] = Course{Id: "040064", Name: "VO Kostenrechnung"}
courses[35] = Course{Id: "040375", Name: "VO Buchhaltung"}
courses[36] = Course{Id: "040628", Name: "VO Bilanzierung"}
courses[37] = Course{Id: "053020", Name: "VU Advanced Software Engineering"}
courses[38] = Course{Id: "040058", Name: "VO ABWL Unternehmensführung I"}
courses[39] = Course{Id: "040026", Name: "UE ABWL Unternehmensführung II"}
courses[40] = Course{Id: "040110", Name: "UK Urheber-, Patent-, Marken-, Muster- und Ausstattungsrecht"}
courses[41] = Course{Id: "040145", Name: "UK Special Topics in Produktion/Logistik/SCM: Transportlogistik"}
courses[42] = Course{Id: "040284", Name: "UK Rechtsfragen des E-Commerce"}
courses[43] = Course{Id: "040601", Name: "VO Steuerrecht"}
courses[44] = Course{Id: "040602", Name: "UK Steuerrecht"}
courses[45] = Course{Id: "040603", Name: "UK Grundzüge des Unternehmenssteuerrechts"}
courses[46] = Course{Id: "040604", Name: "UK Sonderfragen des Unternehmenssteuerrechts"}
courses[47] = Course{Id: "040712", Name: "UK Unternehmensrecht"}
courses[48] = Course{Id: "040715", Name: "UK Rechtliche Rahmenbedingungen der Rechnungslegung"}
courses[49] = Course{Id: "050074", Name: "VU Innovationsmanagement"}
courses[50] = Course{Id: "052512", Name: "VU Interoperability"}
courses[51] = Course{Id: "053220", Name: "VU Metamodellierung"}
courses[52] = Course{Id: "052411", Name: "VU Business Intelligence I"}
courses[53] = Course{Id: "053225", Name: "VU Semantische Technologien"}
courses[54] = Course{Id: "180087", Name: "EV-L STEOP: Einführung in die theoretische Philosophie"}
courses[55] = Course{Id: "180001", Name: "EV-L STEOP: Einführung in die praktische Philosophie"}
courses[56] = Course{Id: "180004", Name: "LPS Foucault: Sexualität und Wahrheit"}
courses[57] = Course{Id: "180005", Name: "LPS Texte zur philosophischen Praxis"}
courses[58] = Course{Id: "180006", Name: "LPS Emmanuel Levinas: Die Spur des Anderen"}
courses[59] = Course{Id: "180007", Name: "LPS Das Erhabene: Pseudo-Longinus, Kant, Schiller, Adorno, Lyotard"}
courses[60] = Course{Id: "180009", Name: "LPS Einführung in die Strukturanthropologie"}
courses[61] = Course{Id: "180010", Name: "LPS Kants Grundlegung zur Metaphysik der Sitten"}
courses[62] = Course{Id: "180011", Name: "LPS Kant - Kritik der reinen Vernunft"}
courses[63] = Course{Id: "180089", Name: "IK Einführung in das wissenschaftliche Arbeiten in der Philosophie"}
courses[64] = Course{Id: "180090", Name: "GV Griechische Terminologie"}
courses[65] = Course{Id: "180091", Name: "LPS Der junge Spinoza"}
courses[66] = Course{Id: "180131", Name: "LPS Texte zur Medien- und Technikphilosophie"}
courses[67] = Course{Id: "180168", Name: "LPS Aristoteles - Nikomachische Ethik"}
courses[68] = Course{Id: "180012", Name: "VO-GKL Vorlesung Grundkurs Logik"}
courses[69] = Course{Id: "180013", Name: "VO Einführung in die Sprachphilosophie"}
courses[70] = Course{Id: "180183", Name: "UE-GKL Übung Logik"}
courses[71] = Course{Id: "180184", Name: "IK Rhetorik und Argumentationstheorie"}
courses[72] = Course{Id: "180014", Name: "VO-L Geschichte der Philosophie I (Antike)"}
courses[73] = Course{Id: "180085", Name: "VO-L Geschichte der Philosophie II"}
courses[74] = Course{Id: "180017", Name: "PS Edmund Husserl - Cartesianische Meditationen"}
courses[75] = Course{Id: "180018", Name: "PS Eine Einführung zu Plotin"}
courses[76] = Course{Id: "180019", Name: "PS Die phänomenologische Reduktion"}
courses[77] = Course{Id: "180020", Name: "SE Platon Staat"}
courses[78] = Course{Id: "180021", Name: "SE Kausalität und Freiheit"}
courses[79] = Course{Id: "180022", Name: "VO-L Kant: Kritik der Urteilskraft"}
courses[80] = Course{Id: "180023", Name: "SE Construction of social reality"}
courses[81] = Course{Id: "180024", Name: "VO Einführung in die Wissenschaftsphilosophie"}
courses[82] = Course{Id: "180088", Name: "VO Erkenntnistheorie"}
courses[83] = Course{Id: "180093", Name: "SE Philosophical Fragments in Contemporary Philosophy"}
courses[84] = Course{Id: "180118", Name: "PS Ontology and Metaphysics"}
courses[85] = Course{Id: "180125", Name: "Grenzen des Glaubens bei Kant und Hegel"}
courses[86] = Course{Id: "180128", Name: "PS Aristoteles: Metaphysik"}
courses[87] = Course{Id: "180132", Name: "SE The Phenomenology of Habits"}
courses[98] = Course{Id: "030297", Name: "KU Transformations of Constitutionalism "}
courses[99] = Course{Id: "180025", Name: "PS Das Erscheinen der Moral"}
courses[90] = Course{Id: "180026", Name: "PS Naturethik"}
courses[91] = Course{Id: "180027", Name: "SE Rechtfertigungen des Unrechts"}
courses[92] = Course{Id: "180028", Name: "VO-L Moral im Kontext des liberalen Rechtsstaats"}
courses[93] = Course{Id: "180029", Name: "VO-L Praktische Rationalität"}
courses[94] = Course{Id: "180031", Name: "VO-L Philosophische Grundlagen des Antisemitismus"}
courses[95] = Course{Id: "180037", Name: "SE Verantwortung in Technik, Medien und Wissenschaft"}
courses[96] = Course{Id: "180122", Name: "SE Ethik, Menschenwürde und Menschenrechte in der Psychoanalyse"}
courses[97] = Course{Id: "180130", Name: "VO-L 1933/2017. Rechtsextremismus. Elite. Volksbildung"}
courses[98] = Course{Id: "180133", Name: "PS Ethik der Philosophie - Ethik der Religion"}
courses[99] = Course{Id: "180134", Name: "SE Philosophie und Phänomenologie der Gewalt"}
courses[100] = Course{Id: "180136", Name: "PS Armut und Philosophie"}
courses[101] = Course{Id: "180138", Name: "VO Freiheit, Gleichheit, Macht "}
courses[102] = Course{Id: "180157", Name: "PS Ethical Questions: Theory and Practice"}
courses[103] = Course{Id: "180158", Name: "SE On Social Agency"}
courses[104] = Course{Id: "180165", Name: "VO Grundlagen der angewandten Ethik"}
courses[105] = Course{Id: "180032", Name: "SE Podcasts philosophisch"}
courses[106] = Course{Id: "180033", Name: "SE Umwelt, Milieu, Nische - Medien vor den Medien"}
courses[107] = Course{Id: "180034", Name: "SE Money as Medium and Technology"}
courses[108] = Course{Id: "180035", Name: "VO-L Media, Technology, and Romanticism"}
courses[109] = Course{Id: "180036", Name: "SE Hannah Arendts kritische Reflexionen über Technik und Kultur"}
courses[110] = Course{Id: "180037", Name: "SE Verantwortung in Technik, Medien und Wissenschaft "}
courses[111] = Course{Id: "180002", Name: "SE Dimensionen interkultureller Philosophie"}
courses[112] = Course{Id: "180038", Name: "SE Der Begriff Kultur und die Frage der kulturellen Identität"}
courses[113] = Course{Id: "180040", Name: "VO-L Die Philosophie des Tantrismus"}
courses[114] = Course{Id: "180123", Name: "VO-L Zur Genealogie der Humanistischen Philosophie des 20. Jahrhu"}
courses[115] = Course{Id: "180124", Name: "SE Methodologische Grundlagen der Komparativen Philosophie"}
courses[116] = Course{Id: "180126", Name: "SE Philosophie in Lateinamerika 2 ( Peru, Ecuador, Bolivien)"}
courses[117] = Course{Id: "180042", Name: "SE Grundfragen der Daseinsanalyse"}
courses[118] = Course{Id: "180047", Name: "SE Hegels Wissenschaft der Logik: Sein, Wesen, Begriff"}
courses[119] = Course{Id: "180052", Name: "VO Current Developments in Philosophy of Mind and Cognitive Science"}
courses[120] = Course{Id: "180053", Name: "Subjekt, Modus, Kraft"}
courses[121] = Course{Id: "180057", Name: "Emotions, Feelings, and Moods"}
courses[122] = Course{Id: "180058", Name: "SE Philosophische Anthropologie als historisch-genetische Entwicklungstheorie"}
courses[123] = Course{Id: "010025", Name: "VO Klassiker der Religionsphilosophie"}
courses[124] = Course{Id: "140121", Name: "VM1 / VM7 - Coloniality under De_Construction "}
courses[125] = Course{Id: "180049", Name: "SE Räume"}
courses[126] = Course{Id: "180051", Name: "VO-L Soziale Implikationen der Religion"}
courses[127] = Course{Id: "180078", Name: "VO-L Elemente der philosophischen Aesthetik"}
courses[128] = Course{Id: "180124", Name: "SE Methodologische Grundlagen der Komparativen Philosophie"}
courses[129] = Course{Id: "180126", Name: "SE Philosophie in Lateinamerika 2 ( Peru, Ecuador, Bolivien)"}
courses[130] = Course{Id: "180042", Name: "SE Grundfragen der Daseinsanalyse"}
courses[131] = Course{Id: "180047", Name: "SE Hegels Wissenschaft der Logik: Sein, Wesen, Begriff "}
courses[132] = Course{Id: "180052", Name: "Current Developments in Philosophy of Mind and Cognitive Science"}
courses[133] = Course{Id: "180053", Name: "Subjekt, Modus, Kraft"}
courses[134] = Course{Id: "180057", Name: "Emotions, Feelings, and Moods"}
courses[135] = Course{Id: "180058", Name: "SE Philosophische Anthropologie als historisch-genetische Entwicklungstheorie"}
courses[136] = Course{Id: "010025", Name: "VO Klassiker der Religionsphilosophie"}
courses[137] = Course{Id: "140121", Name: "VM1 / VM7 - Coloniality under De_Construction"}
courses[138] = Course{Id: "180049", Name: "VO Klassiker der Religionsphilosophie"}
courses[139] = Course{Id: "140121", Name: "VM1 / VM7 - Coloniality under De_Construction"}
courses[140] = Course{Id: "180049", Name: "SE Räume"}
courses[141] = Course{Id: "180051", Name: "VO-L Soziale Implikationen der Religion"}
courses[142] = Course{Id: "180078", Name: "VO-L Elemente der philosophischen Aesthetik"}
courses[143] = Course{Id: "180124", Name: "SE Methodologische Grundlagen der Komparativen Philosophie"}
courses[144] = Course{Id: "180140", Name: "SE Auch Klio dichtet"}
courses[145] = Course{Id: "180176", Name: "SE Philosophie der Literatur"}
courses[146] = Course{Id: "010096", Name: "SE Diskursethik und lateinamerikanische Ethik der Befreiung"}
courses[147] = Course{Id: "040096", Name: "UK SOLV - Commons - neue alte Wirtschaftsform (BA)"}
courses[148] = Course{Id: "180037", Name: "SE Verantwortung in Technik, Medien und Wissenschaft"}
courses[149] = Course{Id: "180042", Name: "SE Grundfragen der Daseinsanalyse"}
courses[150] = Course{Id: "030502", Name: "Transformations of Constitutionalism"}
courses[151] = Course{Id: "180055", Name: "SE Der Gebrauch der Lüste"}
courses[152] = Course{Id: "180056", Name: "SE Kapitalismus und Schizophrenie "}
courses[153] = Course{Id: "180077", Name: "SE Der Mensch auf der Flucht"}
courses[154] = Course{Id: "180121", Name: "VO-L Philosophie der Wissenschaft und die Frage der Forschung in der Psychoanalyse"}
courses[155] = Course{Id: "180138", Name: "VO Freiheit, Gleichheit, Macht"}
courses[156] = Course{Id: "180142", Name: "SE Bodies that matter"}
courses[157] = Course{Id: "010090", Name: "VO Aufbaukurs Philosophische Anthropologie"}
courses[158] = Course{Id: "180055", Name: "SE Der Gebrauch der Lüste"}
courses[159] = Course{Id: "180057", Name: "Emotions, Feelings, and Moods"}
courses[160] = Course{Id: "180058", Name: "SE Philosophische Anthropologie als historisch-genetische Entwicklungstheorie"}
courses[161] = Course{Id: "180071", Name: "SE Grundlegungen der Praxis"}
courses[162] = Course{Id: "180080", Name: "SE Was ist Leben?"}
courses[163] = Course{Id: "180129", Name: "VO Das biologische Menschenbild und seine philosophischen Konsequenzen"}
courses[164] = Course{Id: "180047", Name: "SE Hegels Wissenschaft der Logik: Sein, Wesen, Begriff "}
courses[165] = Course{Id: "180140", Name: "SE Auch Klio dichtet"}
courses[166] = Course{Id: "180143", Name: "SE Wissenschaftslogik"}
courses[167] = Course{Id: "180145", Name: "SE Wissenschaftstheorie"}
courses[168] = Course{Id: "040686", Name: "Repetitorium: Grundzüge der ABWL"}
courses[169] = Course{Id: "040641", Name: "VO Soziologie für WirtschaftswissenschafterInnen"}
courses[170] = Course{Id: "040002", Name: "VO ABWL Finanzwirtschaft I"}
courses[171] = Course{Id: "040003", Name: "ABWL Finanzwirtschaft II"}
courses[172] = Course{Id: "041016", Name: "UK Repetitorium zu VO ABWL Finanzwirtschaft I"}
courses[173] = Course{Id: "040138", Name: "UK ABWL Marketing II"}
courses[174] = Course{Id: "040192", Name: "ABWL Marketing I"}
courses[175] = Course{Id: "040240", Name: "ABWL Marketing II"}
courses[176] = Course{Id: "040026", Name: "ABWL Unternehmensführung II"}
courses[177] = Course{Id: "040058", Name: "VO ABWL Unternehmensführung I"}
courses[178] = Course{Id: "040136", Name: "VO ABWL Produktion und Logistik I"}
courses[179] = Course{Id: "040137", Name: "UK ABWL Produktion und Logistik II"}
courses[180] = Course{Id: "040600", Name: "VK ABWL Innovations- und Technologiemanagement"}
courses[181] = Course{Id: "040064", Name: "VO Kostenrechnung"}
courses[182] = Course{Id: "040628", Name: "VO Bilanzierung"}
courses[183] = Course{Id: "040005", Name: "UK Entscheidungstheorie"}
courses[184] = Course{Id: "040184", Name: "Mikroökonomie"}
courses[185] = Course{Id: "040373", Name: "VO Grundzüge des Rechts"}
courses[186] = Course{Id: "040037", Name: "VO Gesellschaftsrecht"}
courses[187] = Course{Id: "040712", Name: "UK Unternehmensrecht"}
courses[188] = Course{Id: "040601", Name: "VO Steuerrecht"}
courses[189] = Course{Id: "040602", Name: "UK Steuerrecht"}
courses[190] = Course{Id: "040006", Name: "VO Informationstechnologie"}
courses[191] = Course{Id: "040130", Name: "Business English II"}
courses[192] = Course{Id: "040131", Name: "Business English I"}
courses[193] = Course{Id: "040291", Name: "Special Topics in Banking and Finance: International Finance"}
courses[194] = Course{Id: "040883", Name: "Special Topics in Banking and Finance: Risk and Insurance"}
courses[195] = Course{Id: "040071", Name: "UK Marketing Instruments B"}
courses[196] = Course{Id: "040331", Name: "UK Organisationsstrukturen und Prozesse"}
courses[197] = Course{Id: "040608", Name: "Personalmanagement"}
courses[198] = Course{Id: "040162", Name: "Special Topics in Produktion/Logistik/SCM: Planspiel zu Produktionsmanagem"}
courses[199] = Course{Id: "040489", Name: "UK Special Topics in Produktion/Logistik/SCM: Produktionsmanagement"}
courses[200] = Course{Id: "040050", Name: "Group Accounting"}
courses[201] = Course{Id: "040627", Name: "UK Sonderprobleme der Bilanzierung"}
courses[202] = Course{Id: "040631", Name: "UK Rechnungslegung nach IFRS (I)"}
courses[203] = Course{Id: "040640", Name: "UK Konzernrechnungslegung"}
courses[204] = Course{Id: "040715", Name: "UK Rechtliche Rahmenbedingungen der Rechnungslegung"}
courses[205] = Course{Id: "040071", Name: "UK Marketing Instruments B"}
courses[206] = Course{Id: "040291", Name: "Special Topics in Banking and Finance: International Finance"}
courses[207] = Course{Id: "040331", Name: "UK Organisationsstrukturen und Prozesse"}
courses[208] = Course{Id: "040608", Name: "Personalmanagement"}
courses[209] = Course{Id: "040627", Name: "UK Sonderprobleme der Bilanzierung"}
courses[210] = Course{Id: "040715", Name: "UK Rechtliche Rahmenbedingungen der Rechnungslegung"}
courses[211] = Course{Id: "040883", Name: "Special Topics in Banking and Finance: Risk and Insurance"}
courses[212] = Course{Id: "040093", Name: "UK Wettbewerbs- und Kartellrecht"}
courses[213] = Course{Id: "040102", Name: "UK Privatstiftungsrecht"}
courses[214] = Course{Id: "040103", Name: "UK Genossenschafts- und Vereinsrecht"}
courses[215] = Course{Id: "040107", Name: "UK Konzernrecht"}
courses[216] = Course{Id: "040110", Name: "UK Urheber-, Patent-, Marken-, Muster- und Ausstattungsrecht"}
courses[217] = Course{Id: "040111", Name: "UK Umgründungsrecht"}
courses[218] = Course{Id: "040126", Name: "UK Besonderes Wirtschaftsrecht"}
courses[219] = Course{Id: "040241", Name: "UK Internationales Privat- und Europarecht"}
courses[220] = Course{Id: "040284", Name: "UK Rechtsfragen des E-Commerce"}
courses[221] = Course{Id: "040588", Name: "UK Internationale Steuerplanung"}
courses[222] = Course{Id: "040589", Name: "UK Sonderfragen der Besteuerung von Finanzinstrumenten"}
courses[223] = Course{Id: "040590", Name: "UK Arbeits- und Sozialrecht"}
courses[224] = Course{Id: "040603", Name: "UK Grundzüge des Unternehmenssteuerrechts"}
courses[225] = Course{Id: "040604", Name: "UK Sonderfragen des Unternehmenssteuerrechts"}
courses[226] = Course{Id: "040050", Name: "Group Accounting"}
courses[227] = Course{Id: "040627", Name: "UK Sonderprobleme der Bilanzierung"}
courses[228] = Course{Id: "040631", Name: "UK Rechnungslegung nach IFRS (I)"}
courses[229] = Course{Id: "040715", Name: "UK Rechtliche Rahmenbedingungen der Rechnungslegung"}
courses[230] = Course{Id: "040261", Name: "UK Integrierte betriebliche Informationssysteme 2"}
courses[231] = Course{Id: "040007", Name: "Unternehmensmodellierung 2"}
courses[232] = Course{Id: "040032", Name: "FK BW VE: Wahlfach E-Marketing (Teil 2)"}
courses[233] = Course{Id: "040060", Name: "FK BW VE: Wahlfach E-Marketing (Teil 1)"}
courses[234] = Course{Id: "040149", Name: "UK Lineare Multivariate Statistik"}
courses[235] = Course{Id: "040690", Name: "UK Erweiterungen des Linearen Modells"}
courses[236] = Course{Id: "040109", Name: "UK Lineare Modelle 2"}
courses[237] = Course{Id: "040460", Name: "STEOP: Wahrscheinlichkeitsrechnung (UE)"}
courses[238] = Course{Id: "040414", Name: "PR Statistisches Consulting"}
courses[239] = Course{Id: "040971", Name: "UK Computational Statistics"}
courses[240] = Course{Id: "040754", Name: "UK Kooperative Spiele (BA)"}
courses[241] = Course{Id: "040106", Name: "Entscheidungs- und Spieltheorie (BA)"}
courses[242] = Course{Id: "040099", Name: "UK Grundzüge der Finanzwissenschaft (BA)"}
courses[243] = Course{Id: "040159", Name: "UK Steuern (BA)"}
courses[244] = Course{Id: "040120", Name: "UE WI-SPR: Business English Vorbereitungskurs"}
courses[245] = Course{Id: "040176", Name: "KU Repetitorium: Kostenrechnung"}
courses[246] = Course{Id: "040940", Name: "VK Wirtschaftsdeutsch"}
courses[247] = Course{Id: "140233", Name: "STEOP - Einführung in die afrikanische Geschichtswissenschaft"}
courses[248] = Course{Id: "140239", Name: "STEOP - Einführung in die afrikanische Literaturwissenschaft"}
courses[249] = Course{Id: "140240", Name: "STEOP: Einführung in die afrikanische Sprachwissenschaft"}
courses[250] = Course{Id: "140241", Name: "VO+UE Einführung in das wissenschaftliche Arbeiten"}
courses[251] = Course{Id: "140244", Name: "VO+UE Vertiefung afrikanische Sprachwissenschaft"}
courses[252] = Course{Id: "140309", Name: "VO+UE Vertiefung afrikanische Literaturwissenschaft"}
courses[253] = Course{Id: "140321", Name: "VO+UE Vertiefung afrikanische Geschichtswissenschaft"}
courses[254] = Course{Id: "140207", Name: "VO+UE Bambara: Übungen 2"}
courses[255] = Course{Id: "140235", Name: "VO+UE Bambara: Grammatik 2"}
courses[256] = Course{Id: "140293", Name: "VO+UE Hausa: Grammatik 2"}
courses[257] = Course{Id: "140311", Name: "VO+UE Hausa: Übungen 2"}
courses[258] = Course{Id: "140271", Name: "VO+UE Swahili: Übungen 2 / Kiswahili: Mazoezi 2"}
courses[259] = Course{Id: "140313", Name: "VO+UE Swahili: Grammatik 2 / Kiswahili: Sarufi 2"}
courses[260] = Course{Id: "140314", Name: "SK Bambara Grammatik 4"}
courses[261] = Course{Id: "140315", Name: "SK Bambara: Texte 2"}
courses[262] = Course{Id: "140327", Name: "SK Bambara: Konversation 2"}
courses[263] = Course{Id: "140295", Name: "Hausa: Texts 2"}
courses[264] = Course{Id: "140319", Name: "SK Hausa: Grammatik 4"}
courses[265] = Course{Id: "140329", Name: "Hausa: Conversation 2"}
courses[266] = Course{Id: "140182", Name: "SK Swahili: Konversation 2 / Kiswahili: Namna ya kuzungumza 2"}
courses[267] = Course{Id: "140333", Name: "SK Swahili: Grammatik 4 / Kiswahili: Sarufi 4"}
courses[268] = Course{Id: "140334", Name: "SK Swahili: Texte 2"}
courses[269] = Course{Id: "140232", Name: "SE Bachelorseminar: Vorkoloniale und koloniale Geschichte Afrikas"}
courses[270] = Course{Id: "140234", Name: "PS Proseminar Afrikanische Geschichte: Zwischen Dorf und Welt"}
courses[271] = Course{Id: "140237", Name: "VO Geschichte Afrikas vom 16. bis 18. Jahrhundert"}
courses[272] = Course{Id: "140338", Name: "VO Wissenschaftsgeschichte der Afrikawissenschaften"}
courses[273] = Course{Id: "140345", Name: "VO Geschichte Westafrikas 2"}
courses[274] = Course{Id: "140346", Name: "VO Geschichte Nordafrikas 2"}
courses[275] = Course{Id: "140385", Name: "VO Geschichte Südafrikas: Bilder im Kopf und vor der Kamera"}
courses[276] = Course{Id: "140223", Name: "African Poetry Workshop"}
courses[277] = Course{Id: "140265", Name: "Contemporary African Women's Writing and African Feminism"}
courses[278] = Course{Id: "140338", Name: "VO Wissenschaftsgeschichte der Afrikawissenschaften"}
courses[279] = Course{Id: "140340", Name: "SE Bachelorseminar: Afrikanische Sprach- und Literaturwissenschaft"}
courses[280] = Course{Id: "140341", Name: "PS Widerstand im Film: Politisches Kino in Afrika und der Diaspora"}
courses[281] = Course{Id: "140382", Name: "VO Literaturen Afrikas vor dem 20. Jahrhundert"}
courses[282] = Course{Id: "140216", Name: "The Africa Diaspora in China: Topics in Sociocultural Linguistics and Beyond"}
courses[283] = Course{Id: "140251", Name: "VO Konfliktfähigkeit und Konfliktpotenzial: Sprachliche Stilmittel im Umgang"}
courses[284] = Course{Id: "140260", Name: "PS Proseminar: Soziolinguistik"}
courses[285] = Course{Id: "140338", Name: "VO Wissenschaftsgeschichte der Afrikawissenschaften"}
courses[286] = Course{Id: "140339", Name: "Grammatical Description: Comparative Studies of Language Structures"}
courses[287] = Course{Id: "140340", Name: "SE Bachelorseminar: Afrikanische Sprach- und Literaturwissenschaft"}
courses[288] = Course{Id: "140386", Name: "VO Einführung in die Morphologie afrikanischer Sprachen"}
courses[289] = Course{Id: "140209", Name: "KU Wissenschaftliche Texte: Gestalten und Präsentieren"}
courses[290] = Course{Id: "140220", Name: "Aktuelle Forschungsgebiete in den Afrikawissenschaften"}
courses[291] = Course{Id: "140371", Name: "SE Master Coaching Seminar"}
courses[292] = Course{Id: "140226", Name: "Structure of a L3 (Dagaare and the Mabia Languages)"}
courses[293] = Course{Id: "140230", Name: "KU Struktur einer L 2/2 (Fulfulde)"}
courses[294] = Course{Id: "140247", Name: "VO Sprache und Geschichte: Nil-Sahara-Sahel"}
courses[295] = Course{Id: "140317", Name: "Syntactic Theory and Analysis: LFG"}
courses[296] = Course{Id: "140225", Name: "Advances in African Linguistic and Literary Studies"}
courses[297] = Course{Id: "140278", Name: "VO Sprachdenkmäler von den Kanarischen Inseln"}
courses[298] = Course{Id: "140371", Name: "SE Master Coaching Seminar"}
courses[299] = Course{Id: "140356", Name: "KU Mündliche Literaturen Afrikas in Geschichte und Gegenwart"}
courses[300] = Course{Id: "140357", Name: "African Literature and the Anthropology of the Other"}
courses[301] = Course{Id: "140225", Name: "Advances in African Linguistic and Literary Studies"}
courses[302] = Course{Id: "140264", Name: "VO Afrikanische Kulturen und Literaturen in der Diaspora"}
courses[303] = Course{Id: "140359", Name: "SE Performativität und Agency im afrikanischen Film"}
courses[304] = Course{Id: "140371", Name: "SE Master Coaching Seminar"}
courses[305] = Course{Id: "140121", Name: "VM1 / VM7 - Coloniality under De_Construction"}
courses[306] = Course{Id: "140254", Name: "SE Nation, Gender, Race und Religion"}
courses[307] = Course{Id: "140255", Name: "KU Oral History in den historisch perspektivierten Afrikawissenschaften"}
courses[308] = Course{Id: "140283", Name: "VO Das Reich der schwarzen Pharaonen: Kusch und die antike Welt"}
courses[309] = Course{Id: "140316", Name: "VO Postkoloniale, post-säkulare und Geschlechtertheorie"}
courses[310] = Course{Id: "140373", Name: "VO+UE VM5 / VM1 - Afrikanischer Nationalismus, Sozialismus und Marxismus"}
courses[311] = Course{Id: "140254", Name: "SE Nation, Gender, Race und Religion"}
courses[312] = Course{Id: "140283", Name: "VO Das Reich der schwarzen Pharaonen: Kusch und die antike Welt"}
courses[313] = Course{Id: "140316", Name: "VO Postkoloniale, post-säkulare und Geschlechtertheorie"}
courses[314] = Course{Id: "140209", Name: "KU Wissenschaftliche Texte: Gestalten und Präsentieren"}
courses[315] = Course{Id: "140225", Name: "Advances in African Linguistic and Literary Studies"}
courses[316] = Course{Id: "140247", Name: "VO Sprache und Geschichte: Nil-Sahara-Sahel"}
courses[317] = Course{Id: "140278", Name: "VO Sprachdenkmäler von den Kanarischen Inseln"}
courses[318] = Course{Id: "140317", Name: "Syntactic Theory and Analysis: LFG"}
courses[319] = Course{Id: "140371", Name: "SE Master Coaching Seminar"}
courses[320] = Course{Id: "140121", Name: "VM1 / VM7 - Coloniality under De_Construction"}
courses[321] = Course{Id: "140254", Name: "SE Nation, Gender, Race und Religion"}
courses[322] = Course{Id: "140255", Name: "KU Oral History in den historisch perspektivierten Afrikawissenschaften"}
courses[323] = Course{Id: "140283", Name: "VO Das Reich der schwarzen Pharaonen: Kusch und die antike Welt"}
courses[324] = Course{Id: "140316", Name: "VO Postkoloniale, post-säkulare und Geschlechtertheorie"}
courses[325] = Course{Id: "140371", Name: "SE Master Coaching Seminar"}
courses[326] = Course{Id: "140373", Name: "VO+UE VM5 / VM1 - Afrikanischer Nationalismus, Sozialismus und Marxismus "}
courses[327] = Course{Id: "140223", Name: "African Poetry Workshop"}
courses[328] = Course{Id: "140237", Name: "VO Geschichte Afrikas vom 16. bis 18. Jahrhundert"}
courses[329] = Course{Id: "140265", Name: "Contemporary African Women's Writing and African Feminism"}
courses[330] = Course{Id: "140345", Name: "VO Geschichte Westafrikas 2"}
courses[331] = Course{Id: "140346", Name: "VO Geschichte Nordafrikas 2"}
courses[332] = Course{Id: "140382", Name: "VO Literaturen Afrikas vor dem 20. Jahrhundert"}
courses[333] = Course{Id: "140385", Name: "VO Geschichte Südafrikas: Bilder im Kopf und vor der Kamera"}
courses[334] = Course{Id: "160101", Name: "VO STEOP: Einführung in die Allgemeine Sprachwissenschaft"}
courses[335] = Course{Id: "160102", Name: "VO STEOP: Einführung in die Angewandte Sprachwissenschaft"}
courses[336] = Course{Id: "160103", Name: "VO STEOP: Einführung in die Phonetik und Phonologie"}
courses[337] = Course{Id: "160106", Name: "PS Grundlagen der Allgemeinen Sprachwissenschaft"}
courses[338] = Course{Id: "160107", Name: "VO Einführung in die Semantik und Pragmatik"}
courses[339] = Course{Id: "160108", Name: "PS Grundlagen der Angewandten Sprachwissenschaft"}
courses[340] = Course{Id: "160109", Name: "VO Einführung in die Statistik"}
courses[341] = Course{Id: "160110", Name: "PS Einführung in die Transkription und Gesprächsanalyse"}
courses[342] = Course{Id: "160113", Name: "VO Einführung in die Indogermanistik"}
courses[343] = Course{Id: "160114", Name: "PS Grundlagen der Indogermanistik"}
courses[344] = Course{Id: "160116", Name: "VO Einführung in die Sprachlehr-/-lernforschung"}
courses[345] = Course{Id: "160158", Name: "PS Weiterführende LV aus Diskursanalyse und Soziolinguistik I"}
courses[346] = Course{Id: "160140", Name: "Proseminar aus Grammatiktheorie und kognitiver Sprachwissenschaft"}
courses[347] = Course{Id: "160129", Name: "PS Besondere Probleme der Sprachenpolitik"}
courses[348] = Course{Id: "140317", Name: "Syntactic Theory and Analysis: LFG"}
courses[349] = Course{Id: "140278", Name: "VO Sprachdenkmäler von den Kanarischen Inseln"}
courses[350] = Course{Id: "140247", Name: "VO Sprache und Geschichte: Nil-Sahara-Sahel"}
courses[351] = Course{Id: "160117", Name: "PS Methoden der Angewandten Sprachwissenschaft"}
courses[352] = Course{Id: "160101", Name: "VO STEOP: Einführung in die Allgemeine Sprachwissenschaft"}
courses[353] = Course{Id: "160103", Name: "VO STEOP: Einführung in die Phonetik und Phonologie"}
courses[354] = Course{Id: "160102", Name: "VO STEOP: Einführung in die Angewandte Sprachwissenschaft"}
courses[355] = Course{Id: "160105", Name: "PS Proseminar Grundlagen der Allgemeinen Sprachwissenschaft B"}
courses[356] = Course{Id: "160106", Name: "PS Grundlagen der Allgemeinen Sprachwissenschaft"}
courses[357] = Course{Id: "160108", Name: "PS Grundlagen der Angewandten Sprachwissenschaft"}
courses[358] = Course{Id: "160113", Name: "VO Einführung in die Indogermanistik"}
courses[359] = Course{Id: "160114", Name: "PS Grundlagen der Indogermanistik"}
courses[360] = Course{Id: "160107", Name: "VO Einführung in die Semantik und Pragmatik"}
courses[361] = Course{Id: "160116", Name: "VO Einführung in die Sprachlehr-/-lernforschung"}
courses[362] = Course{Id: "160120", Name: "SE BA-Seminar aus Grammatiktheorie und kognitiver Linguistik"}
courses[363] = Course{Id: "160122", Name: "SE Empirisches Seminar aus Diskursanalyse und Soziolinguistik I"}
courses[364] = Course{Id: "160123", Name: "SE BA-Seminar aus Psycho- oder Patholinguistik"}
courses[365] = Course{Id: "160124", Name: "SE Empirisches SE aus Sprachlehrforschung, Sprachlernforschung"}
courses[366] = Course{Id: "160110", Name: "PS Einführung in die Transkription und Gesprächsanalyse"}
courses[367] = Course{Id: "160117", Name: "PS Methoden der Angewandten Sprachwissenschaft"}
courses[368] = Course{Id: "160110", Name: "PS Einführung in die Transkription und Gesprächsanalyse"}
courses[369] = Course{Id: "160121", Name: "Psycho- oder Patholinguistisches PS"}
courses[370] = Course{Id: "160123", Name: "SE BA-Seminar aus Psycho- oder Patholinguistik"}
courses[371] = Course{Id: "160126", Name: "PR Wissenschaftliches Praktikum 1"}
courses[372] = Course{Id: "160125", Name: "Einführung in die Semantik"}
courses[373] = Course{Id: "160127", Name: "BA-Seminar aus Grammatiktheorie und kognitiver Linguistik"}
courses[374] = Course{Id: "160128", Name: "PS Grammatiktheoretisches PS "}

	return courses
}
