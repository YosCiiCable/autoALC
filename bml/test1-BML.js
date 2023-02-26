javascript: (
	BtnLogin.click();
	LbtSubCourseLink_1.click(); 
	GoToStUnitList_Click('PWH_L03');

	ShowLearnPage('PWH_L03_U001-1','U001','09', '', 'UNIT001','3', '&STCnt1=');

	document.querySelector(".ui-button-text-icon-secondary").click();

	// ShowLearnPage('PWH_L03_U002-1','U002','09', '', 'UNIT002','3', '&STCnt1=');
	// ShowLearnPage('PWH_L03_U00*-1','U00*','09', '', 'UNIT00*','3', '&STCnt1=');
)();
