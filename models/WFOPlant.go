package models

/**
A rather unimportant comment, but nontheless somewhat interesting if you are not familiar with taxonomy in botany. the terms 'specific epithet' and 'infraspecific epithet' are what is commonly often refered to as species and subspecies. the infraspecific epithet is usually refered to as 'variant'. An example is 'Acanthocalycium klimpelianum var. macranthum' where Acontholacyium is the genus, klimpelianum is the specific epithet, var. short for varietas, signifying that macranthum refers to the infraspecific epithet.
**/

type WFOPlant struct {
	TaxonID				string	`db:"taxonID"			json:"taxonID"`			
	ScientificNameID		string	`db:"scientificNameID"		json:"scientificNameID"`	
	LocalID				string	`db:"localID"			json:"localID"`			
	ScientificName			string	`db:"scientificName"		json:"scientificName"`		
	TaxonRank			string	`db:"taxonRank"			json:"taxonRank"`		
	ParentNameUsageID		string	`db:"parentNameUsageID"		json:"parentNameUsageID"`	
	ScientificNameAuthorship	string	`db:"scientificNameAuthorship"	json:"scientificNameAuthorship"`
	Family				string	`db:"family"			json:"family"`			
	Subfamily			string	`db:"subfamily"			json:"subfamily"`		
	Tribe				string	`db:"tribe"			json:"tribe"`			
	Subtribe			string	`db:"subtribe"			json:"subtribe"`		
	Genus				string	`db:"genus"			json:"genus"`
	Subgenus			string	`db:"subgenus"			json:"subgenus"`		
	SpecificEpithet			string	`db:"specificEpithet"		json:"specificEpithet"`		
	InfraspecificEpithet		string	`db:"infraspecificEpithet"	json:"infraspecificEpithet"`	
	VerbatimTaxonRank		string	`db:"verbatimTaxonRank"		json:"verbatimTaxonRank"`	
	NomenclaturalStatus		string	`db:"nomenclaturalStatus"	json:"nomenclaturalStatus"`	
	NamePublishedIn			string	`db:"namePublishedIn"		json:"namePublishedIn"`		
	TaxonomicStatus			string	`db:"taxonomicStatus"		json:"taxonomicStatus"`		
	AcceptedNameUsageID		string	`db:"acceptedNameUsageID"	json:"acceptedNameUsageID"`	
	OriginalNameUsageID		string	`db:"originalNameUsageID"	json:"originalNameUsageID"`	
	NameAccordingToID		string	`db:"nameAccordingToID"		json:"nameAccordingToID"`	
	TaxonRemarks			string	`db:"taxonRemarks"		json:"taxonRemarks"`		
	Created				string	`db:"created"			json:"created"`
	Modified			string	`db:"modified"			json:"modified"`
	References			string	`db:"references"		json:"references"`		
	Source				string	`db:"source"			json:"source"`			
	MajorGroup			string	`db:"majorGroup"		json:"majorGroup"`		
	TplID				string	`db:"tplID"			json:"tplID"`			
}
