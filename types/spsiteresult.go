package types

type SPSitesResult struct {
	D struct {
		Query struct {
			Metadata struct {
				Type string `json:"type"`
			} `json:"__metadata"`
			Elapsedtime        int `json:"ElapsedTime"`
			Primaryqueryresult struct {
				Metadata struct {
					Type string `json:"type"`
				} `json:"__metadata"`
				Customresults struct {
					Metadata struct {
						Type string `json:"type"`
					} `json:"__metadata"`
					Results []interface{} `json:"results"`
				} `json:"CustomResults"`
				Queryid           string      `json:"QueryId"`
				Queryruleid       string      `json:"QueryRuleId"`
				Refinementresults interface{} `json:"RefinementResults"`
				Relevantresults   struct {
					Metadata struct {
						Type string `json:"type"`
					} `json:"__metadata"`
					Grouptemplateid interface{} `json:"GroupTemplateId"`
					Itemtemplateid  interface{} `json:"ItemTemplateId"`
					Properties      struct {
						Metadata struct {
							Type string `json:"type"`
						} `json:"__metadata"`
						Results []struct {
							Key       string `json:"Key"`
							Value     string `json:"Value"`
							Valuetype string `json:"ValueType"`
						} `json:"results"`
					} `json:"Properties"`
					Resulttitle    interface{} `json:"ResultTitle"`
					Resulttitleurl interface{} `json:"ResultTitleUrl"`
					Rowcount       int         `json:"RowCount"`
					Table          struct {
						Metadata struct {
							Type string `json:"type"`
						} `json:"__metadata"`
						Rows struct {
							Results []struct {
								Metadata struct {
									Type string `json:"type"`
								} `json:"__metadata"`
								Cells struct {
									Results []struct {
										Metadata struct {
											Type string `json:"type"`
										} `json:"__metadata"`
										Key       string `json:"Key"`
										Value     string `json:"Value"`
										Valuetype string `json:"ValueType"`
									} `json:"results"`
								} `json:"Cells"`
							} `json:"results"`
						} `json:"Rows"`
					} `json:"Table"`
					Totalrows                    int `json:"TotalRows"`
					Totalrowsincludingduplicates int `json:"TotalRowsIncludingDuplicates"`
				} `json:"RelevantResults"`
				Specialtermresults interface{} `json:"SpecialTermResults"`
			} `json:"PrimaryQueryResult"`
			Properties struct {
				Metadata struct {
					Type string `json:"type"`
				} `json:"__metadata"`
				Results []struct {
					Key       string `json:"Key"`
					Value     string `json:"Value"`
					Valuetype string `json:"ValueType"`
				} `json:"results"`
			} `json:"Properties"`
			Secondaryqueryresults struct {
				Metadata struct {
					Type string `json:"type"`
				} `json:"__metadata"`
				Results []interface{} `json:"results"`
			} `json:"SecondaryQueryResults"`
			Spellingsuggestion interface{} `json:"SpellingSuggestion"`
			Triggeredrules     struct {
				Metadata struct {
					Type string `json:"type"`
				} `json:"__metadata"`
				Results []interface{} `json:"results"`
			} `json:"TriggeredRules"`
		} `json:"query"`
	} `json:"d"`
}
