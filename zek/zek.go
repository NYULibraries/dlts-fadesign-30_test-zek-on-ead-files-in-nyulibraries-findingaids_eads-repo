package zek

// Based on https://github.com/NYULibraries/fadesign-30_test-zek-on-ead-files-in-nyulibraries-findingaids_eads-repo/blob/112c8ac853544961c09019375ce807b317985343/zek-types.go

import (
	"encoding/xml"
)

// Ead was generated 2020-02-06 18:13:07
type EAD struct {
	XMLName         xml.Name `xml:"ead"`
	Text            string   `xml:",chardata"`
	SchemaLocation  string   `xml:"schemaLocation,attr"`
	Xsi             string   `xml:"xsi,attr"`
	Ns2             string   `xml:"ns2,attr"`
	Xmlns           string   `xml:"xmlns,attr"`
	Relatedencoding string   `xml:"relatedencoding,attr"`
	Eadheader       struct {
		Text               string `xml:",chardata"`
		Countryencoding    string `xml:"countryencoding,attr"`
		Dateencoding       string `xml:"dateencoding,attr"`
		Findaidstatus      string `xml:"findaidstatus,attr"`
		Langencoding       string `xml:"langencoding,attr"`
		Repositoryencoding string `xml:"repositoryencoding,attr"`
		Audience           string `xml:"audience,attr"`
		ID                 string `xml:"id,attr"`
		Eadid              struct {
			Text           string `xml:",chardata"`
			Countrycode    string `xml:"countrycode,attr"`
			Mainagencycode string `xml:"mainagencycode,attr"`
			URL            string `xml:"url,attr"`
			Identifier     string `xml:"identifier,attr"`
		} `xml:"eadid"`
		Filedesc struct {
			Text      string `xml:",chardata"`
			Titlestmt struct {
				Text        string `xml:",chardata"`
				Titleproper []struct {
					Text           string `xml:",chardata"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Type           string `xml:"type,attr"`
					Emph           []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Num  []string `xml:"num"`
					Lb   []string `xml:"lb"`
					Date []struct {
						Text     string `xml:",chardata"`
						Calendar string `xml:"calendar,attr"`
						Era      string `xml:"era,attr"`
						Normal   string `xml:"normal,attr"`
						Type     string `xml:"type,attr"`
						Lb       string `xml:"lb"`
					} `xml:"date"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"titleproper"`
				Author struct {
					Text           string `xml:",chardata"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Lb             string `xml:"lb"`
				} `xml:"author"`
				Sponsor struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"sponsor"`
				Subtitle string `xml:"subtitle"`
			} `xml:"titlestmt"`
			Editionstmt struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"`
					Date struct {
						Text     string `xml:",chardata"`
						Calendar string `xml:"calendar,attr"`
						Era      string `xml:"era,attr"`
					} `xml:"date"`
				} `xml:"p"`
			} `xml:"editionstmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher struct {
					Text           string `xml:",chardata"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"publisher"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text   string   `xml:",chardata"`
						Lb     []string `xml:"lb"`
						Extptr struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extptr"`
					} `xml:"addressline"`
				} `xml:"address"`
				Date string `xml:"date"`
				P    struct {
					Text string `xml:",chardata"`
					Date struct {
						Text           string `xml:",chardata"`
						Encodinganalog string `xml:"encodinganalog,attr"`
						Normal         string `xml:"normal,attr"`
					} `xml:"date"`
					Address struct {
						Text        string `xml:",chardata"`
						Addressline string `xml:"addressline"`
					} `xml:"address"`
					Extref struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"p"`
			} `xml:"publicationstmt"`
			Notestmt struct {
				Text string `xml:",chardata"`
				Note struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"note"`
			} `xml:"notestmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text           string `xml:",chardata"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Date           struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"date"`
			} `xml:"creation"`
			Langusage struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text           string `xml:",chardata"`
					Langcode       string `xml:"langcode,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Scriptcode     string `xml:"scriptcode,attr"`
				} `xml:"language"`
				Lb string `xml:"lb"`
			} `xml:"langusage"`
			Descrules string `xml:"descrules"`
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
				Text           string `xml:",chardata"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Date           struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"date"`
				Item struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"change"`
		} `xml:"revisiondesc"`
	} `xml:"eadheader"`
	Archdesc struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
		Type  string `xml:"type,attr"`
		Did   struct {
			Text         string `xml:",chardata"`
			ID           string `xml:"id,attr"`
			Langmaterial []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Label    string `xml:"label,attr"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"langmaterial"`
			Repository struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Corpname       string `xml:"corpname"`
				Address        struct {
					Text        string `xml:",chardata"`
					Addressline string `xml:"addressline"`
				} `xml:"address"`
			} `xml:"repository"`
			Unittitle struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Emph           []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Lb       []string `xml:"lb"`
				Unitdate struct {
					Text           string `xml:",chardata"`
					Type           string `xml:"type,attr"`
					Normal         string `xml:"normal,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"unitdate"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"unittitle"`
			Origination []struct {
				Text           string `xml:",chardata"`
				Audience       string `xml:"audience,attr"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Persname       struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Role           string `xml:"role,attr"`
					Rules          string `xml:"rules,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"persname"`
				Corpname struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Rules          string `xml:"rules,attr"`
					Source         string `xml:"source,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"corpname"`
				Famname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"famname"`
			} `xml:"origination"`
			Unitid []struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Countrycode    string `xml:"countrycode,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Repositorycode string `xml:"repositorycode,attr"`
				Audience       string `xml:"audience,attr"`
				Identifier     string `xml:"identifier,attr"`
				Type           string `xml:"type,attr"`
			} `xml:"unitid"`
			Physdesc []struct {
				Text           string `xml:",chardata"`
				Altrender      string `xml:"altrender,attr"`
				ID             string `xml:"id,attr"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Audience       string `xml:"audience,attr"`
				Extent         []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
				Physfacet string `xml:"physfacet"`
			} `xml:"physdesc"`
			Unitdate []struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Abstract []struct {
				Text           string   `xml:",chardata"`
				ID             string   `xml:"id,attr"`
				Label          string   `xml:"label,attr"`
				Encodinganalog string   `xml:"encodinganalog,attr"`
				Lb             []string `xml:"lb"`
				Emph           []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Title  string `xml:"title,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"abstract"`
			Physloc []struct {
				Text           string   `xml:",chardata"`
				ID             string   `xml:"id,attr"`
				Label          string   `xml:"label,attr"`
				Audience       string   `xml:"audience,attr"`
				Encodinganalog string   `xml:"encodinganalog,attr"`
				Lb             []string `xml:"lb"`
				Extref         struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Href    string `xml:"href,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"physloc"`
			Container []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Type      string `xml:"type,attr"`
				Parent    string `xml:"parent,attr"`
			} `xml:"container"`
			Head         string `xml:"head"`
			Materialspec struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Label string `xml:"label,attr"`
			} `xml:"materialspec"`
		} `xml:"did"`
		Custodhist []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Lb     []string `xml:"lb"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
				Date struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
			} `xml:"p"`
		} `xml:"custodhist"`
		Accessrestrict []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
			Legalstatus string `xml:"legalstatus"`
		} `xml:"accessrestrict"`
		Userestrict []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
				} `xml:"extref"`
				Lb      []string `xml:"lb"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Corpname string `xml:"corpname"`
			} `xml:"p"`
		} `xml:"userestrict"`
		Prefercite []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string `xml:",chardata"`
						Lb   string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Bioghist []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Extref    struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
				} `xml:"emph"`
				List []struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Bibref struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"bibref"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
					} `xml:"item"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"list"`
				Lb     []string `xml:"lb"`
				Extref []struct {
					Text     string `xml:",chardata"`
					Href     string `xml:"href,attr"`
					Type     string `xml:"type,attr"`
					Actuate  string `xml:"actuate,attr"`
					Show     string `xml:"show,attr"`
					Title    string `xml:"title,attr"`
					Linktype string `xml:"linktype,attr"`
					Archref  string `xml:"archref"`
					Emph     struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"extref"`
				Name     []string `xml:"name"`
				Corpname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"corpname"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				Geogname  []string `xml:"geogname"`
				Abbr      string   `xml:"abbr"`
				Chronlist struct {
					Text      string `xml:",chardata"`
					Chronitem []struct {
						Text string `xml:",chardata"`
						Date struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Event struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"event"`
						Eventgrp struct {
							Text  string `xml:",chardata"`
							Event []struct {
								Text  string `xml:",chardata"`
								Title struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"event"`
						} `xml:"eventgrp"`
					} `xml:"chronitem"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"chronlist"`
				Occupation []string `xml:"occupation"`
				Table      struct {
					Text   string `xml:",chardata"`
					Tgroup struct {
						Text    string `xml:",chardata"`
						Cols    string `xml:"cols,attr"`
						Colspec []struct {
							Text     string `xml:",chardata"`
							Colwidth string `xml:"colwidth,attr"`
						} `xml:"colspec"`
						Tbody struct {
							Text string `xml:",chardata"`
							Row  []struct {
								Text  string `xml:",chardata"`
								Entry []struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"entry"`
							} `xml:"row"`
						} `xml:"tbody"`
					} `xml:"tgroup"`
				} `xml:"table"`
				Blockquote struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"blockquote"`
				Persname []struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"persname"`
				Date []struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Subject []string `xml:"subject"`
				Bibref  struct {
					Text  string `xml:",chardata"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"bibref"`
				Archref struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Type string `xml:"type,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Name  string `xml:"name"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
					} `xml:"extref"`
				} `xml:"item"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
			Chronlist []struct {
				Text     string `xml:",chardata"`
				Audience string `xml:"audience,attr"`
				Head     struct {
					Text  string `xml:",chardata"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"head"`
				Chronitem []struct {
					Text     string `xml:",chardata"`
					Date     string `xml:"date"`
					Eventgrp struct {
						Text  string `xml:",chardata"`
						Event []struct {
							Text  string `xml:",chardata"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"title"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb     []string `xml:"lb"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"event"`
					} `xml:"eventgrp"`
					Event struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb string `xml:"lb"`
					} `xml:"event"`
				} `xml:"chronitem"`
			} `xml:"chronlist"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Extref []struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
			Lb []string `xml:"lb"`
		} `xml:"bioghist"`
		Scopecontent []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Emph      struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				List []struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb      []string `xml:"lb"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
					} `xml:"item"`
					Head string `xml:"head"`
				} `xml:"list"`
				Lb       []string `xml:"lb"`
				Corpname []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"corpname"`
				Name     []string `xml:"name"`
				Geogname string   `xml:"geogname"`
				Extref   []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Occupation string `xml:"occupation"`
				Archref    []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Date struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Genreform []string `xml:"genreform"`
				I         string   `xml:"i"`
				Persname  string   `xml:"persname"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text     string   `xml:",chardata"`
					Corpname []string `xml:"corpname"`
					Lb       []string `xml:"lb"`
					Title    struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Archref []struct {
						Text   string `xml:",chardata"`
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"archref"`
				} `xml:"item"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  string `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
			Extref struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
			Chronlist struct {
				Text      string `xml:",chardata"`
				Head      string `xml:"head"`
				Chronitem []struct {
					Text  string `xml:",chardata"`
					Date  string `xml:"date"`
					Event string `xml:"event"`
				} `xml:"chronitem"`
			} `xml:"chronlist"`
		} `xml:"scopecontent"`
		Arrangement []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Title  string `xml:"title,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Lb []string `xml:"lb"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				List struct {
					Text       string `xml:",chardata"`
					Numeration string `xml:"numeration,attr"`
					Type       string `xml:"type,attr"`
					Item       []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"item"`
				} `xml:"list"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Audience   string `xml:"audience,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Title  struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
					Lb    []string `xml:"lb"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
					Ref struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"ref"`
				} `xml:"item"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
		} `xml:"arrangement"`
		Relatedmaterial []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Title   string `xml:"title,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Target  string `xml:"target,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
					Emph struct {
						Text   string `xml:",chardata"`
						Title  string `xml:"title,attr"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"extref"`
				Lb      []string `xml:"lb"`
				Archref []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Emph    struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"extref"`
				} `xml:"archref"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Emph []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Extref    struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
				} `xml:"emph"`
				List struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
					} `xml:"item"`
				} `xml:"list"`
				Ref struct {
					Text  string `xml:",chardata"`
					Title string `xml:"title,attr"`
				} `xml:"ref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Audience   string `xml:"audience,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text    string `xml:",chardata"`
					Archref struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Extref struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"archref"`
					List struct {
						Text string   `xml:",chardata"`
						Item []string `xml:"item"`
					} `xml:"list"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
					Title []struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"item"`
			} `xml:"list"`
			Extref []struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Actuate string `xml:"actuate,attr"`
				Emph    struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"extref"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Lb []string `xml:"lb"`
		} `xml:"relatedmaterial"`
		Separatedmaterial []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				List []struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Lb []string `xml:"lb"`
					} `xml:"item"`
					Head string `xml:"head"`
				} `xml:"list"`
				Archref struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Type string `xml:"type,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
				Head string `xml:"head"`
			} `xml:"list"`
		} `xml:"separatedmaterial"`
		Phystech []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"phystech"`
		Controlaccess struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Geogname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"geogname"`
			Subject []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"subject"`
			Genreform []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"genreform"`
			Persname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Audience       string `xml:"audience,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"persname"`
			Corpname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Audience       string `xml:"audience,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"corpname"`
			Title []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"title"`
			Famname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Audience       string `xml:"audience,attr"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"famname"`
			Occupation []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
			Head          string `xml:"head"`
			Controlaccess []struct {
				Text     string `xml:",chardata"`
				Head     string `xml:"head"`
				Persname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"persname"`
				Corpname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"corpname"`
				Subject []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"subject"`
				Geogname []struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"geogname"`
				Genreform []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"genreform"`
				Famname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"famname"`
			} `xml:"controlaccess"`
			Function struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"function"`
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Type string `xml:"type,attr"`
			C    C      `xml:"c"`
			C01 []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Did   struct {
					Text      string `xml:",chardata"`
					Unittitle struct {
						Text           string `xml:",chardata"`
						Encodinganalog string `xml:"encodinganalog,attr"`
						Emph           []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"unittitle"`
					Container []struct {
						Text   string `xml:",chardata"`
						ID     string `xml:"id,attr"`
						Type   string `xml:"type,attr"`
						Label  string `xml:"label,attr"`
						Parent string `xml:"parent,attr"`
					} `xml:"container"`
					Unitdate []struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						Language struct {
							Text     string `xml:",chardata"`
							Langcode string `xml:"langcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Altrender string `xml:"altrender,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
						Dimensions struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"dimensions"`
					} `xml:"physdesc"`
					Unitid []struct {
						Text       string `xml:",chardata"`
						Audience   string `xml:"audience,attr"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Audience string `xml:"audience,attr"`
						Persname struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
						} `xml:"persname"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
							Role   string `xml:"role,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
				} `xml:"did"`
				C02 []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string   `xml:",chardata"`
							Lb   []string `xml:"lb"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"emph"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"unittitle"`
						Container []struct {
							Text   string `xml:",chardata"`
							ID     string `xml:"id,attr"`
							Type   string `xml:"type,attr"`
							Label  string `xml:"label,attr"`
							Parent string `xml:"parent,attr"`
						} `xml:"container"`
						Unitdate []struct {
							Text   string `xml:",chardata"`
							Normal string `xml:"normal,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"unitdate"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Altrender string `xml:"altrender,attr"`
							Extent    []struct {
								Text      string `xml:",chardata"`
								Unit      string `xml:"unit,attr"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"extent"`
							Dimensions struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"dimensions"`
						} `xml:"physdesc"`
						Origination []struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Audience string `xml:"audience,attr"`
							Persname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
						} `xml:"origination"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							Language struct {
								Text     string `xml:",chardata"`
								Langcode string `xml:"langcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
						Unitid []struct {
							Text       string `xml:",chardata"`
							Audience   string `xml:"audience,attr"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
					} `xml:"did"`
					Odd []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"odd"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Controlaccess struct {
						Text     string `xml:",chardata"`
						Corpname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"corpname"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"persname"`
						Subject []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"subject"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Geogname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Occupation []struct {
							Text           string `xml:",chardata"`
							Authfilenumber string `xml:"authfilenumber,attr"`
							Source         string `xml:"source,attr"`
						} `xml:"occupation"`
						Title struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"title"`
					} `xml:"controlaccess"`
					Phystech struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"phystech"`
					C03 []struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"id,attr"`
						Level string `xml:"level,attr"`
						Did   struct {
							Text      string `xml:",chardata"`
							Container []struct {
								Text   string `xml:",chardata"`
								Type   string `xml:"type,attr"`
								ID     string `xml:"id,attr"`
								Label  string `xml:"label,attr"`
								Parent string `xml:"parent,attr"`
							} `xml:"container"`
							Unittitle struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
								Lb []string `xml:"lb"`
							} `xml:"unittitle"`
							Unitdate []struct {
								Text     string `xml:",chardata"`
								Normal   string `xml:"normal,attr"`
								Era      string `xml:"era,attr"`
								Calendar string `xml:"calendar,attr"`
								Type     string `xml:"type,attr"`
							} `xml:"unitdate"`
							Langmaterial struct {
								Text     string `xml:",chardata"`
								Language struct {
									Text     string `xml:",chardata"`
									Langcode string `xml:"langcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
						} `xml:"did"`
						Phystech struct {
							Text string   `xml:",chardata"`
							P    []string `xml:"p"`
						} `xml:"phystech"`
						C04 []struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Level string `xml:"level,attr"`
							Did   struct {
								Text      string `xml:",chardata"`
								Unittitle struct {
									Text  string `xml:",chardata"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
									Lb   []string `xml:"lb"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"unittitle"`
								Container []struct {
									Text   string `xml:",chardata"`
									ID     string `xml:"id,attr"`
									Type   string `xml:"type,attr"`
									Label  string `xml:"label,attr"`
									Parent string `xml:"parent,attr"`
								} `xml:"container"`
								Unitdate struct {
									Text   string `xml:",chardata"`
									Normal string `xml:"normal,attr"`
									Type   string `xml:"type,attr"`
								} `xml:"unitdate"`
							} `xml:"did"`
							Odd struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    struct {
									Text  string   `xml:",chardata"`
									Lb    []string `xml:"lb"`
									Title []struct {
										Text   string   `xml:",chardata"`
										Render string   `xml:"render,attr"`
										Lb     []string `xml:"lb"`
									} `xml:"title"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"odd"`
							C05 []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Level string `xml:"level,attr"`
								Did   struct {
									Text      string `xml:",chardata"`
									Unittitle struct {
										Text  string `xml:",chardata"`
										Lb    string `xml:"lb"`
										Title struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"title"`
									} `xml:"unittitle"`
									Container []struct {
										Text   string `xml:",chardata"`
										ID     string `xml:"id,attr"`
										Type   string `xml:"type,attr"`
										Label  string `xml:"label,attr"`
										Parent string `xml:"parent,attr"`
									} `xml:"container"`
									Unitdate struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
								} `xml:"did"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    struct {
										Text  string `xml:",chardata"`
										Title struct {
											Text   string   `xml:",chardata"`
											Render string   `xml:"render,attr"`
											Lb     []string `xml:"lb"`
										} `xml:"title"`
										Lb []string `xml:"lb"`
									} `xml:"p"`
								} `xml:"odd"`
							} `xml:"c05"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"scopecontent"`
						} `xml:"c04"`
						Odd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text  string   `xml:",chardata"`
								Lb    []string `xml:"lb"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"odd"`
						Scopecontent struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text  string `xml:",chardata"`
								Title struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"p"`
						} `xml:"scopecontent"`
					} `xml:"c03"`
					Accessrestrict struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accessrestrict"`
					Relatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Show    string `xml:"show,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"relatedmaterial"`
					Dao []struct {
						Text     string `xml:",chardata"`
						Actuate  string `xml:"actuate,attr"`
						Show     string `xml:"show,attr"`
						Title    string `xml:"title,attr"`
						Role     string `xml:"role,attr"`
						Href     string `xml:"href,attr"`
						Type     string `xml:"type,attr"`
						Audience string `xml:"audience,attr"`
						Daodesc  struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Arrangement struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"arrangement"`
					Processinfo struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"processinfo"`
					Altformavail struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"altformavail"`
					Custodhist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"custodhist"`
					Fileplan struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"fileplan"`
				} `xml:"c02"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Emph   []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"emph"`
						Lb    []string `xml:"lb"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
					Head string `xml:"head"`
				} `xml:"scopecontent"`
				Arrangement struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"arrangement"`
				Dao []struct {
					Text     string `xml:",chardata"`
					Actuate  string `xml:"actuate,attr"`
					Show     string `xml:"show,attr"`
					Title    string `xml:"title,attr"`
					Role     string `xml:"role,attr"`
					Href     string `xml:"href,attr"`
					Audience string `xml:"audience,attr"`
					Type     string `xml:"type,attr"`
					Daodesc  struct {
						Text string `xml:",chardata"`
						P    string `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Odd struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"p"`
				} `xml:"odd"`
				Controlaccess struct {
					Text     string `xml:",chardata"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"corpname"`
					Persname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"persname"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Subject []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"subject"`
					Geogname []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"geogname"`
					Occupation []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"occupation"`
					Title struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"title"`
				} `xml:"controlaccess"`
				Bioghist struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"bioghist"`
			} `xml:"c01"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"dsc"`
		Acqinfo []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"acqinfo"`
		Processinfo []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     string `xml:"head"`
			P        []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
				} `xml:"extref"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb   []string `xml:"lb"`
				I    []string `xml:"i"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"processinfo"`
		Appraisal []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Rend   string `xml:"rend,attr"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"appraisal"`
		Otherfindaid struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
				} `xml:"extref"`
				Archref struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
			} `xml:"p"`
		} `xml:"otherfindaid"`
		Accruals struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accruals"`
		Bibliography struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     string `xml:"head"`
			P        []struct {
				Text   string `xml:",chardata"`
				Bibref []struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
				} `xml:"bibref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
			Bibref []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
			} `xml:"bibref"`
		} `xml:"bibliography"`
		Odd []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				List struct {
					Text string   `xml:",chardata"`
					Item []string `xml:"item"`
				} `xml:"list"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text     string   `xml:",chardata"`
					Persname []string `xml:"persname"`
					Corpname string   `xml:"corpname"`
					Title    struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Render string `xml:"render,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Title  struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"odd"`
		Altformavail struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"altformavail"`
		Originalsloc struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"originalsloc"`
		Dao struct {
			Text    string `xml:",chardata"`
			Actuate string `xml:"actuate,attr"`
			Show    string `xml:"show,attr"`
			Title   string `xml:"title,attr"`
			Role    string `xml:"role,attr"`
			Href    string `xml:"href,attr"`
			Daodesc struct {
				Text string `xml:",chardata"`
				P    string `xml:"p"`
			} `xml:"daodesc"`
		} `xml:"dao"`
		Index struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Head       string `xml:"head"`
			P          string `xml:"p"`
			Indexentry []struct {
				Text  string `xml:",chardata"`
				Title string `xml:"title"`
				Ref   string `xml:"ref"`
			} `xml:"indexentry"`
		} `xml:"index"`
	} `xml:"archdesc"`
}

type C []struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"id,attr"`
	Level      string `xml:"level,attr"`
	Otherlevel string `xml:"otherlevel,attr"`
	Audience   string `xml:"audience,attr"`
	Did        struct {
		Text      string `xml:",chardata"`
		Unittitle struct {
			Text string `xml:",chardata"`
			Emph []struct {
				Text      string `xml:",chardata"`
				Render    string `xml:"render,attr"`
				AttrTitle string `xml:"title,attr"`
				Title     struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"emph"`
			Title []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"title"`
			Lb   []string `xml:"lb"`
			Name struct {
				Text string `xml:",chardata"`
				Role string `xml:"role,attr"`
			} `xml:"name"`
			Persname string `xml:"persname"`
			Corpname string `xml:"corpname"`
		} `xml:"unittitle"`
		Unitid []struct {
			Text       string `xml:",chardata"`
			Audience   string `xml:"audience,attr"`
			Identifier string `xml:"identifier,attr"`
			Type       string `xml:"type,attr"`
		} `xml:"unitid"`
		Unitdate []struct {
			Text      string `xml:",chardata"`
			Normal    string `xml:"normal,attr"`
			Type      string `xml:"type,attr"`
			Certainty string `xml:"certainty,attr"`
		} `xml:"unitdate"`
		Container []struct {
			Text      string `xml:",chardata"`
			Altrender string `xml:"altrender,attr"`
			ID        string `xml:"id,attr"`
			Label     string `xml:"label,attr"`
			Type      string `xml:"type,attr"`
			Parent    string `xml:"parent,attr"`
		} `xml:"container"`
		Physdesc []struct {
			Text      string `xml:",chardata"`
			Altrender string `xml:"altrender,attr"`
			ID        string `xml:"id,attr"`
			Label     string `xml:"label,attr"`
			Audience  string `xml:"audience,attr"`
			Extent    []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
			} `xml:"extent"`
			Physfacet struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"physfacet"`
			Emph struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Dimensions struct {
				Text     string `xml:",chardata"`
				Audience string `xml:"audience,attr"`
			} `xml:"dimensions"`
		} `xml:"physdesc"`
		Origination []struct {
			Text     string `xml:",chardata"`
			Audience string `xml:"audience,attr"`
			Label    string `xml:"label,attr"`
			Persname struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Rules  string `xml:"rules,attr"`
				Role   string `xml:"role,attr"`
			} `xml:"persname"`
			Corpname struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Rules  string `xml:"rules,attr"`
				Role   string `xml:"role,attr"`
			} `xml:"corpname"`
		} `xml:"origination"`
		Langmaterial struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Language struct {
				Text     string `xml:",chardata"`
				Langcode string `xml:"langcode,attr"`
			} `xml:"language"`
		} `xml:"langmaterial"`
		Abstract struct {
			Text  string   `xml:",chardata"`
			ID    string   `xml:"id,attr"`
			Label string   `xml:"label,attr"`
			Lb    []string `xml:"lb"`
		} `xml:"abstract"`
		Physloc struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id,attr"`
			Label   string `xml:"label,attr"`
			Archref struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"archref"`
		} `xml:"physloc"`
		Materialspec struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Label string `xml:"label,attr"`
		} `xml:"materialspec"`
	} `xml:"did"`
	Scopecontent []struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text string `xml:",chardata"`
			Emph []struct {
				Text      string `xml:",chardata"`
				Render    string `xml:"render,attr"`
				Renderr   string `xml:"renderr,attr"`
				AttrTitle string `xml:"title,attr"`
				Title     struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"emph"`
			Lb    []string `xml:"lb"`
			Title []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
				Type   string `xml:"type,attr"`
				Emph   struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"title"`
			Extref []struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Actuate string `xml:"actuate,attr"`
				Emph    struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"extref"`
			Corpname string `xml:"corpname"`
			Date     []struct {
				Text     string `xml:",chardata"`
				Calendar string `xml:"calendar,attr"`
				Era      string `xml:"era,attr"`
			} `xml:"date"`
			Subject    []string `xml:"subject"`
			Genreform  []string `xml:"genreform"`
			Occupation string   `xml:"occupation"`
			Archref    struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"archref"`
			List struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"list"`
			I   string `xml:"i"`
			Ref []struct {
				Text   string `xml:",chardata"`
				Target string `xml:"target,attr"`
			} `xml:"ref"`
			Blockquote struct {
				Text string `xml:",chardata"`
				P    struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Title   []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"extref"`
				} `xml:"p"`
			} `xml:"blockquote"`
		} `xml:"p"`
		List struct {
			Text       string `xml:",chardata"`
			Numeration string `xml:"numeration,attr"`
			Type       string `xml:"type,attr"`
			Head       string `xml:"head"`
			Item       []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"item"`
			Defitem struct {
				Text  string `xml:",chardata"`
				Label string `xml:"label"`
				Item  string `xml:"item"`
			} `xml:"defitem"`
		} `xml:"list"`
		Emph []struct {
			Text   string `xml:",chardata"`
			Render string `xml:"render,attr"`
		} `xml:"emph"`
		Lb     []string `xml:"lb"`
		Extref struct {
			Text    string `xml:",chardata"`
			Actuate string `xml:"actuate,attr"`
			Href    string `xml:"href,attr"`
		} `xml:"extref"`
	} `xml:"scopecontent"`
	C []struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"id,attr"`
		Level      string `xml:"level,attr"`
		Otherlevel string `xml:"otherlevel,attr"`
		Audience   string `xml:"audience,attr"`
		Did        struct {
			Text      string `xml:",chardata"`
			Unittitle struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Title  string `xml:"title,attr"`
					Lb     string `xml:"lb"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Name struct {
					Text string `xml:",chardata"`
					Role string `xml:"role,attr"`
				} `xml:"name"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string   `xml:",chardata"`
					Render string   `xml:"render,attr"`
					Type   string   `xml:"type,attr"`
					Lb     []string `xml:"lb"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				I        string `xml:"i"`
				Corpname string `xml:"corpname"`
				Extref   struct {
					Text     string `xml:",chardata"`
					Linktype string `xml:"linktype,attr"`
					Actuate  string `xml:"actuate,attr"`
					Show     string `xml:"show,attr"`
					Href     string `xml:"href,attr"`
				} `xml:"extref"`
				Persname string `xml:"persname"`
			} `xml:"unittitle"`
			Unitid []struct {
				Text       string `xml:",chardata"`
				Audience   string `xml:"audience,attr"`
				Identifier string `xml:"identifier,attr"`
				Type       string `xml:"type,attr"`
			} `xml:"unitid"`
			Unitdate []struct {
				Text      string `xml:",chardata"`
				Normal    string `xml:"normal,attr"`
				Type      string `xml:"type,attr"`
				Certainty string `xml:"certainty,attr"`
			} `xml:"unitdate"`
			Container []struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Type      string `xml:"type,attr"`
				Parent    string `xml:"parent,attr"`
				Altrender string `xml:"altrender,attr"`
			} `xml:"container"`
			Physdesc []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Audience  string `xml:"audience,attr"`
				Extent    []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
				Physfacet struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"physfacet"`
				Dimensions struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
				} `xml:"dimensions"`
			} `xml:"physdesc"`
			Materialspec struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"materialspec"`
			Physloc struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Label string `xml:"label,attr"`
			} `xml:"physloc"`
			Langmaterial struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Label    string `xml:"label,attr"`
				Language struct {
					Text     string `xml:",chardata"`
					Langcode string `xml:"langcode,attr"`
				} `xml:"language"`
			} `xml:"langmaterial"`
			Abstract struct {
				Text  string   `xml:",chardata"`
				ID    string   `xml:"id,attr"`
				Label string   `xml:"label,attr"`
				Lb    []string `xml:"lb"`
			} `xml:"abstract"`
			Origination []struct {
				Text     string `xml:",chardata"`
				Audience string `xml:"audience,attr"`
				Label    string `xml:"label,attr"`
				Persname struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Role           string `xml:"role,attr"`
					Rules          string `xml:"rules,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"persname"`
				Corpname struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Source         string `xml:"source,attr"`
					Rules          string `xml:"rules,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"corpname"`
				Famname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
				} `xml:"famname"`
			} `xml:"origination"`
		} `xml:"did"`
		Phystech []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head struct {
				Text string `xml:",chardata"`
				Lb   string `xml:"lb"`
			} `xml:"head"`
			P []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
			Lb []string `xml:"lb"`
		} `xml:"phystech"`
		C []struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Level      string `xml:"level,attr"`
			Otherlevel string `xml:"otherlevel,attr"`
			Audience   string `xml:"audience,attr"`
			Did        struct {
				Text      string `xml:",chardata"`
				Unittitle struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Emph   string `xml:"emph"`
						Title  struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
					Name  []string `xml:"name"`
					Lb    []string `xml:"lb"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
						Rende  string `xml:"rende,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Archref string `xml:"archref"`
					} `xml:"extref"`
					I        string   `xml:"i"`
					Persname []string `xml:"persname"`
				} `xml:"unittitle"`
				Unitid []struct {
					Text       string `xml:",chardata"`
					Audience   string `xml:"audience,attr"`
					Identifier string `xml:"identifier,attr"`
					Type       string `xml:"type,attr"`
				} `xml:"unitid"`
				Unitdate []struct {
					Text      string `xml:",chardata"`
					Normal    string `xml:"normal,attr"`
					Type      string `xml:"type,attr"`
					Certainty string `xml:"certainty,attr"`
				} `xml:"unitdate"`
				Container []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
					ID        string `xml:"id,attr"`
					Label     string `xml:"label,attr"`
					Type      string `xml:"type,attr"`
					Parent    string `xml:"parent,attr"`
				} `xml:"container"`
				Physdesc []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
					ID        string `xml:"id,attr"`
					Label     string `xml:"label,attr"`
					Audience  string `xml:"audience,attr"`
					Extent    []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						Lb        string `xml:"lb"`
					} `xml:"extent"`
					Physfacet struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"physfacet"`
					Dimensions struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"dimensions"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"physdesc"`
				Origination []struct {
					Text     string `xml:",chardata"`
					Audience string `xml:"audience,attr"`
					Label    string `xml:"label,attr"`
					Persname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
						Rules  string `xml:"rules,attr"`
						Role   string `xml:"role,attr"`
					} `xml:"persname"`
					Corpname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
						Role   string `xml:"role,attr"`
						Rules  string `xml:"rules,attr"`
					} `xml:"corpname"`
				} `xml:"origination"`
				Physloc struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Label string `xml:"label,attr"`
				} `xml:"physloc"`
				Langmaterial struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Label    string `xml:"label,attr"`
					Language struct {
						Text     string `xml:",chardata"`
						Langcode string `xml:"langcode,attr"`
					} `xml:"language"`
				} `xml:"langmaterial"`
				Materialspec struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"materialspec"`
				Abstract struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Label string `xml:"label,attr"`
				} `xml:"abstract"`
			} `xml:"did"`
			Scopecontent []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    []struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
					Lb    []string `xml:"lb"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Extref []struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
					} `xml:"extref"`
					Ref []struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
					} `xml:"ref"`
				} `xml:"p"`
				List struct {
					Text       string `xml:",chardata"`
					Numeration string `xml:"numeration,attr"`
					Type       string `xml:"type,attr"`
					Item       []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
					Head string `xml:"head"`
				} `xml:"list"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"scopecontent"`
			Altformavail struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"emph"`
				} `xml:"p"`
			} `xml:"altformavail"`
			Accessrestrict []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Audience string `xml:"audience,attr"`
				Head     string `xml:"head"`
				P        []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"p"`
				Legalstatus struct {
					Text     string `xml:",chardata"`
					Audience string `xml:"audience,attr"`
				} `xml:"legalstatus"`
			} `xml:"accessrestrict"`
			Phystech []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Lb []string `xml:"lb"`
				} `xml:"p"`
				Lb []string `xml:"lb"`
			} `xml:"phystech"`
			Dao []struct {
				Text    string `xml:",chardata"`
				Actuate string `xml:"actuate,attr"`
				Href    string `xml:"href,attr"`
				Role    string `xml:"role,attr"`
				Show    string `xml:"show,attr"`
				Title   string `xml:"title,attr"`
				Type    string `xml:"type,attr"`
				Daodesc struct {
					Text string `xml:",chardata"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Name  []string `xml:"name"`
						Lb    []string `xml:"lb"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"daodesc"`
			} `xml:"dao"`
			Odd []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Audience string `xml:"audience,attr"`
				Head     string `xml:"head"`
				P        []struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Lb     []string `xml:"lb"`
					Extent string   `xml:"extent"`
					Num    struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"num"`
					Title  string `xml:"title"`
					Extref struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
					Blockquote string `xml:"blockquote"`
				} `xml:"p"`
			} `xml:"odd"`
			C []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				Level      string `xml:"level,attr"`
				Otherlevel string `xml:"otherlevel,attr"`
				Audience   string `xml:"audience,attr"`
				Did        struct {
					Text      string `xml:",chardata"`
					Unittitle struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Title  string `xml:"title,attr"`
						} `xml:"emph"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Lb []string `xml:"lb"`
					} `xml:"unittitle"`
					Unitid []struct {
						Text       string `xml:",chardata"`
						Audience   string `xml:"audience,attr"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Unitdate struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
					Container []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
						Parent    string `xml:"parent,attr"`
					} `xml:"container"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
					} `xml:"physdesc"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Audience string `xml:"audience,attr"`
						Label    string `xml:"label,attr"`
						Persname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"persname"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						Language struct {
							Text     string `xml:",chardata"`
							Langcode string `xml:"langcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
					Physloc struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"physloc"`
				} `xml:"did"`
				Odd []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Extref struct {
							Text    string `xml:",chardata"`
							Type    string `xml:"type,attr"`
							Actuate string `xml:"actuate,attr"`
							Show    string `xml:"show,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"odd"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						I    string `xml:"i"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"scopecontent"`
				Phystech struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"phystech"`
				Accessrestrict struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"accessrestrict"`
				C []struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					Level      string `xml:"level,attr"`
					Otherlevel string `xml:"otherlevel,attr"`
					Did        struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"unittitle"`
						Unitid []struct {
							Text       string `xml:",chardata"`
							Audience   string `xml:"audience,attr"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
						Unitdate struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Normal string `xml:"normal,attr"`
						} `xml:"unitdate"`
						Container []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
						} `xml:"container"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							Language struct {
								Text     string `xml:",chardata"`
								Langcode string `xml:"langcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
					} `xml:"did"`
					Odd struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"odd"`
					Dao []struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Role    string `xml:"role,attr"`
						Show    string `xml:"show,attr"`
						Title   string `xml:"title,attr"`
						Type    string `xml:"type,attr"`
						Daodesc struct {
							Text string   `xml:",chardata"`
							P    []string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Separatedmaterial struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"separatedmaterial"`
				} `xml:"c"`
				Dao []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Role    string `xml:"role,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Daodesc struct {
						Text string `xml:",chardata"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Separatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"separatedmaterial"`
				Controlaccess struct {
					Text     string `xml:",chardata"`
					Geogname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
					Subject struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"subject"`
					Persname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"persname"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
				} `xml:"controlaccess"`
				Bioghist struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"bioghist"`
				Userestrict struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"userestrict"`
				Daogrp struct {
					Text    string `xml:",chardata"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Daodesc struct {
						Text string `xml:",chardata"`
						P    string `xml:"p"`
					} `xml:"daodesc"`
					Daoloc []struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title string `xml:"title,attr"`
						Type  string `xml:"type,attr"`
					} `xml:"daoloc"`
				} `xml:"daogrp"`
				Relatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text   string `xml:",chardata"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"p"`
				} `xml:"relatedmaterial"`
			} `xml:"c"`
			Separatedmaterial struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    struct {
					Text string   `xml:",chardata"`
					Lb   []string `xml:"lb"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Genreform struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Archref struct {
						Text    string `xml:",chardata"`
						Physloc string `xml:"physloc"`
					} `xml:"archref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"p"`
			} `xml:"separatedmaterial"`
			Controlaccess struct {
				Text      string `xml:",chardata"`
				Genreform []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"genreform"`
				Subject []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"subject"`
				Geogname []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"geogname"`
				Persname []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"persname"`
				Corpname []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"corpname"`
				Title []struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"title"`
				Function struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"function"`
				Occupation struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"occupation"`
			} `xml:"controlaccess"`
			Userestrict struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    string `xml:"p"`
			} `xml:"userestrict"`
			Otherfindaid struct {
				Text string   `xml:",chardata"`
				ID   string   `xml:"id,attr"`
				Head string   `xml:"head"`
				P    []string `xml:"p"`
			} `xml:"otherfindaid"`
			Fileplan struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"p"`
			} `xml:"fileplan"`
			Originalsloc struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    string `xml:"p"`
			} `xml:"originalsloc"`
			Arrangement struct {
				Text string   `xml:",chardata"`
				ID   string   `xml:"id,attr"`
				Head string   `xml:"head"`
				P    []string `xml:"p"`
			} `xml:"arrangement"`
			Bioghist struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    []struct {
					Text  string `xml:",chardata"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"p"`
			} `xml:"bioghist"`
			Acqinfo struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    string `xml:"p"`
			} `xml:"acqinfo"`
			Custodhist struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    string `xml:"p"`
			} `xml:"custodhist"`
			Relatedmaterial struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
					Ref struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
					} `xml:"ref"`
				} `xml:"p"`
			} `xml:"relatedmaterial"`
			Processinfo struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"p"`
			} `xml:"processinfo"`
			Accruals struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Head string `xml:"head"`
				P    string `xml:"p"`
			} `xml:"accruals"`
			Appraisal struct {
				Text string   `xml:",chardata"`
				ID   string   `xml:"id,attr"`
				Head string   `xml:"head"`
				P    []string `xml:"p"`
			} `xml:"appraisal"`
		} `xml:"c"`
		Accessrestrict []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"head"`
			P []struct {
				Text string `xml:",chardata"`
				A    struct {
					Text   string `xml:",chardata"`
					Href   string `xml:"href,attr"`
					Target string `xml:"target,attr"`
				} `xml:"a"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"p"`
			Legalstatus struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"legalstatus"`
		} `xml:"accessrestrict"`
		Odd []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb     []string `xml:"lb"`
				Extref struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"odd"`
		Scopecontent []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     struct {
				Text string `xml:",chardata"`
				Lb   string `xml:"lb"`
			} `xml:"head"`
			P []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string   `xml:",chardata"`
					Render string   `xml:"render,attr"`
					Lb     []string `xml:"lb"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
				} `xml:"extref"`
				I       []string `xml:"i"`
				Archref struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Blockquote struct {
					Text string `xml:",chardata"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb []string `xml:"lb"`
					} `xml:"p"`
				} `xml:"blockquote"`
				Chronlist struct {
					Text      string `xml:",chardata"`
					Chronitem []struct {
						Text string `xml:",chardata"`
						Date struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Event string `xml:"event"`
					} `xml:"chronitem"`
				} `xml:"chronlist"`
				Ref []struct {
					Text   string `xml:",chardata"`
					Target string `xml:"target,attr"`
				} `xml:"ref"`
			} `xml:"p"`
			Lb   []string `xml:"lb"`
			List struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Ref  struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
					} `xml:"ref"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"scopecontent"`
		Dao []struct {
			Text     string `xml:",chardata"`
			Actuate  string `xml:"actuate,attr"`
			Href     string `xml:"href,attr"`
			Role     string `xml:"role,attr"`
			Show     string `xml:"show,attr"`
			Title    string `xml:"title,attr"`
			Type     string `xml:"type,attr"`
			Audience string `xml:"audience,attr"`
			Daodesc  struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Name string `xml:"name"`
					Lb   string `xml:"lb"`
				} `xml:"p"`
			} `xml:"daodesc"`
		} `xml:"dao"`
		Separatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"separatedmaterial"`
		Altformavail []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Extref struct {
					Text     string `xml:",chardata"`
					Linktype string `xml:"linktype,attr"`
					Actuate  string `xml:"actuate,attr"`
					Show     string `xml:"show,attr"`
					Href     string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"altformavail"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"arrangement"`
		Fileplan struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"fileplan"`
		Controlaccess struct {
			Text      string `xml:",chardata"`
			Genreform []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"genreform"`
			Subject []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"subject"`
			Persname []struct {
				Text           string `xml:",chardata"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"persname"`
			Occupation []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
			Geogname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"geogname"`
			Corpname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"corpname"`
			Title []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"title"`
			Famname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Role   string `xml:"role,attr"`
			} `xml:"famname"`
		} `xml:"controlaccess"`
		Processinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"processinfo"`
		Appraisal struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"appraisal"`
		Bioghist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"bioghist"`
		Userestrict struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"userestrict"`
		Accruals struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accruals"`
		Originalsloc struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"originalsloc"`
		Custodhist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"custodhist"`
		Relatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Ref struct {
					Text   string `xml:",chardata"`
					Target string `xml:"target,attr"`
				} `xml:"ref"`
			} `xml:"p"`
			Extref struct {
				Text    string `xml:",chardata"`
				Actuate string `xml:"actuate,attr"`
				Href    string `xml:"href,attr"`
			} `xml:"extref"`
		} `xml:"relatedmaterial"`
		Otherfindaid struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     string `xml:"head"`
			P        []struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"otherfindaid"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"prefercite"`
		Index struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Head       string `xml:"head"`
			Indexentry []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name"`
				Subject  string `xml:"subject"`
				Corpname string `xml:"corpname"`
			} `xml:"indexentry"`
		} `xml:"index"`
		Bibliography struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref string `xml:"extref"`
			} `xml:"p"`
		} `xml:"bibliography"`
		Acqinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"acqinfo"`
	} `xml:"c"`
	Arrangement struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text string `xml:",chardata"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Title []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"title"`
			Lb     []string `xml:"lb"`
			Extref struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
			Blockquote struct {
				Text string `xml:",chardata"`
				P    string `xml:"p"`
			} `xml:"blockquote"`
			Ref struct {
				Text   string `xml:",chardata"`
				Target string `xml:"target,attr"`
			} `xml:"ref"`
			P string `xml:"p"`
		} `xml:"p"`
		List struct {
			Text       string   `xml:",chardata"`
			Numeration string   `xml:"numeration,attr"`
			Type       string   `xml:"type,attr"`
			Head       string   `xml:"head"`
			Item       []string `xml:"item"`
			Defitem    []struct {
				Text  string `xml:",chardata"`
				Label string `xml:"label"`
				Item  string `xml:"item"`
			} `xml:"defitem"`
		} `xml:"list"`
	} `xml:"arrangement"`
	Phystech []struct {
		Text     string   `xml:",chardata"`
		ID       string   `xml:"id,attr"`
		Audience string   `xml:"audience,attr"`
		Head     string   `xml:"head"`
		P        []string `xml:"p"`
	} `xml:"phystech"`
	Odd []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"id,attr"`
		Audience string `xml:"audience,attr"`
		Head     string `xml:"head"`
		P        []struct {
			Text string `xml:",chardata"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Lb  []string `xml:"lb"`
			Ref struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Show    string `xml:"show,attr"`
				Actuate string `xml:"actuate,attr"`
			} `xml:"ref"`
			Extref struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
		} `xml:"p"`
	} `xml:"odd"`
	Accessrestrict struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text   string `xml:",chardata"`
			Extref struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Actuate string `xml:"actuate,attr"`
				Show    string `xml:"show,attr"`
			} `xml:"extref"`
			Emph struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
		} `xml:"p"`
	} `xml:"accessrestrict"`
	Otherfindaid struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    string `xml:"p"`
	} `xml:"otherfindaid"`
	Controlaccess struct {
		Text     string `xml:",chardata"`
		Persname []struct {
			Text           string `xml:",chardata"`
			Rules          string `xml:"rules,attr"`
			Source         string `xml:"source,attr"`
			Role           string `xml:"role,attr"`
			Authfilenumber string `xml:"authfilenumber,attr"`
		} `xml:"persname"`
		Corpname []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
			Rules  string `xml:"rules,attr"`
			Role   string `xml:"role,attr"`
		} `xml:"corpname"`
		Subject []struct {
			Text           string `xml:",chardata"`
			Source         string `xml:"source,attr"`
			Authfilenumber string `xml:"authfilenumber,attr"`
		} `xml:"subject"`
		Geogname []struct {
			Text           string `xml:",chardata"`
			Source         string `xml:"source,attr"`
			Authfilenumber string `xml:"authfilenumber,attr"`
		} `xml:"geogname"`
		Genreform []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"genreform"`
		Occupation []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"occupation"`
		Function struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"function"`
		Title []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"title"`
		Famname struct {
			Text   string `xml:",chardata"`
			Rules  string `xml:"rules,attr"`
			Source string `xml:"source,attr"`
		} `xml:"famname"`
	} `xml:"controlaccess"`
	Dao []struct {
		Text     string `xml:",chardata"`
		Actuate  string `xml:"actuate,attr"`
		Href     string `xml:"href,attr"`
		Role     string `xml:"role,attr"`
		Show     string `xml:"show,attr"`
		Title    string `xml:"title,attr"`
		Type     string `xml:"type,attr"`
		Audience string `xml:"audience,attr"`
		Daodesc  struct {
			Text string `xml:",chardata"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"daodesc"`
	} `xml:"dao"`
	Accruals struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    string `xml:"p"`
	} `xml:"accruals"`
	Appraisal struct {
		Text string   `xml:",chardata"`
		ID   string   `xml:"id,attr"`
		Head string   `xml:"head"`
		P    []string `xml:"p"`
	} `xml:"appraisal"`
	Bioghist []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"id,attr"`
		Audience string `xml:"audience,attr"`
		Head     string `xml:"head"`
		P        []struct {
			Text string `xml:",chardata"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Title []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"title"`
			I  []string `xml:"i"`
			Lb []string `xml:"lb"`
		} `xml:"p"`
	} `xml:"bioghist"`
	Processinfo struct {
		Text string   `xml:",chardata"`
		ID   string   `xml:"id,attr"`
		Head string   `xml:"head"`
		P    []string `xml:"p"`
	} `xml:"processinfo"`
	Separatedmaterial struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text string   `xml:",chardata"`
			Lb   []string `xml:"lb"`
			Emph struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
		} `xml:"p"`
	} `xml:"separatedmaterial"`
	Relatedmaterial struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text   string `xml:",chardata"`
			Extref struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Actuate string `xml:"actuate,attr"`
			} `xml:"extref"`
			Emph string `xml:"emph"`
		} `xml:"p"`
		Extref struct {
			Text    string `xml:",chardata"`
			Actuate string `xml:"actuate,attr"`
			Href    string `xml:"href,attr"`
		} `xml:"extref"`
	} `xml:"relatedmaterial"`
	Daogrp struct {
		Text    string `xml:",chardata"`
		Title   string `xml:"title,attr"`
		Type    string `xml:"type,attr"`
		Daodesc struct {
			Text string `xml:",chardata"`
			P    string `xml:"p"`
		} `xml:"daodesc"`
		Daoloc []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Role  string `xml:"role,attr"`
			Title string `xml:"title,attr"`
			Type  string `xml:"type,attr"`
		} `xml:"daoloc"`
	} `xml:"daogrp"`
	Custodhist struct {
		Text string   `xml:",chardata"`
		ID   string   `xml:"id,attr"`
		Head string   `xml:"head"`
		P    []string `xml:"p"`
	} `xml:"custodhist"`
	Userestrict struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    []struct {
			Text string `xml:",chardata"`
			Emph struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Lb []string `xml:"lb"`
		} `xml:"p"`
	} `xml:"userestrict"`
	Prefercite struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    string `xml:"p"`
	} `xml:"prefercite"`
	Acqinfo struct {
		Text string   `xml:",chardata"`
		ID   string   `xml:"id,attr"`
		Head string   `xml:"head"`
		P    []string `xml:"p"`
	} `xml:"acqinfo"`
	Originalsloc struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    string `xml:"p"`
	} `xml:"originalsloc"`
	Altformavail struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Head string `xml:"head"`
		P    struct {
			Text string `xml:",chardata"`
			Emph struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"emph"`
			Extref struct {
				Text    string `xml:",chardata"`
				Actuate string `xml:"actuate,attr"`
				Href    string `xml:"href,attr"`
			} `xml:"extref"`
		} `xml:"p"`
	} `xml:"altformavail"`
}
