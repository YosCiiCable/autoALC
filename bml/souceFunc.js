        var winSt;
        //ENTERキーを押下の処理
        OnClickEnter("BtnBack");
        $(function () {
            $('html, body').animate({ scrollTop: $('.ScrollArea').offset().top }, 800);
            //$('.ScrollArea').scrollTop();
        });

        $('#BtnBack').click(function () {
            // btnLogin の場合は、DoLoginのActionを実行する
            $('#form1').attr('action', '/anetn/Student/StUnitList/GoToTop');
            $('#form1').removeAttr('target');
            $('#form1').submit();

            // サブミットが 2 回発生しないよう、false を返す
            return false;

        });

        
        $('#BtnBackUp').click(function () {
            // btnLogin の場合は、DoLoginのActionを実行する
            $('#form1').attr('action', '/anetn/Student/StUnitList/GoToTop');
            $('#form1').removeAttr('target');
            $('#form1').submit();

            // サブミットが 2 回発生しないよう、false を返す
            return false;

        });
        

        // 学習画面
        function ShowLearnPage(unitId, unitNo, unitDivision, unitType, unitTitle, deviceType, unitTrainingCount, obj) {

            
            var url;

            if(1 == 0)
            {
                url = '/anetn/course/contents.htm';
            }
            else if(1 == 1)
            {
                url = '/anetn/course/contents2.htm';
            }
            



            
            var accountId = 's*****';
            var fullName = encodeURIComponent('漢字');
            var nameKana = encodeURIComponent('ひらがな');
            var nameLatin = encodeURIComponent('Ro-ma zi');
            

            var param = "CCd=1"
                + "&PId=" + '368722'
                
                + "&AId=" + accountId
                + "&FullName=" + fullName
                + "&Kana=" + nameKana
                + "&Latin=" + nameLatin
                
                + "&CId=" + 'PWH'
                + "&VId=" + 'ALC'
                + "&SId=" + 'PWH_L03'
                + "&UId=" + unitId
                + "&UNo=" + unitNo
                + "&TCd=" + unitDivision
                + "&DType=" + deviceType;

            var urlPart = "&SessionId=" + 'dt2mlgoh2x3mi22zd3kp5nkm';

            $('#HidUId').val(unitId);
            $('#HidUNo').val(unitNo);
            $('#HidUCd').val(unitDivision);
            $('#HidDType').val(deviceType);

            $('#DivMessageArea').html("");
            var courseId = 'PWH';
            var vendorId = 'ALC';
            var subCourseId = 'PWH_L03';
            var profileId = '368722';
            if (deviceType == "1") {
                winSt = window.open('', 'SW1', 'fullscreen=yes,scrollbars=yes');
            } else {
                winSt = window.open('', 'SW1', "screenx=0,screeny=0,scrollbars=yes");
            }
            $.ajax({
                url: "/anetn/Student/StUnitList/LinksTitle",
                type: "POST",
                data: {
                    "UnitId": unitId,
                    "UnitTitle": unitTitle,
                    "UnitType": unitType,
                    "CourseId": courseId,
                    "VendorId": vendorId,
                    "SubCourseId": subCourseId,
                    "ProfileId": profileId
                },
                async: false,
                success: function (data) {
                    if (data == null || data.length == 0) {
                        winSt.close();
                        location.href = "/anetn/Error/AJAXError";
                    }
                    var tempArray = (new String(data)).split(",");
                    var isPublic = "0";
                    if (tempArray.length == 3) {

                        param = param + "&AnsFlag=" + tempArray[1] + urlPart;

                        $('#HidAnsFlag').val(tempArray[1]);
                        isPublic = tempArray[0];

                        
                        param = param + "&Qtype=" + tempArray[2];
                        
                    }

                    
                    if(1 == 1)
                    {
                        param += "&DisType=" + '1';
                    }
                    

                    
                    param += unitTrainingCount;
                    

                    //暗号化
                    param = encodeURIComponent(window.btoa(param));

                    url = url + "?queries=" + param;

                    if (unitType == "2") {

                        if (isPublic == "1") {
                            if (deviceType == "1") {
                                winSt.location.href=url;
                                winSt.focus();
                            } else {
                                winSt.location.href=url;
                                winSt.resizeTo(window.screen.availWidth, window.screen.availHeight);
                                winSt.focus();
                            }
                        } else {
                            winSt.close();
                            $('#DivMessageArea').html(tempArray[0]);

                            
                            if (tempArray[1] != null && tempArray[1] != "")
                            {
                                $("#DivTestPublicStartDateTime_" + unitNo).html(tempArray[1]);
                            }

                            if (tempArray[2] != null && tempArray[2] != "")
                            {
                                $("#DivTestPublicEndDateTime_" + unitNo).html(tempArray[2]);
                            }
                            

                            $(obj).parent().html(unitTitle);
                        }
                    } else {
                        if (deviceType == "1") {
                            winSt.location.href=url;
                            winSt.focus();
                        } else {
                            winSt.location.href=url;
                            winSt.resizeTo(window.screen.availWidth, window.screen.availHeight);
                            winSt.focus();
                        }
                    }
                },
                error: function (msg) {
                    winSt.close();
                    location.href = "/anetn/Error/AJAXError";
                }
            });

            // サブミットが 2 回発生しないよう、false を返す
            return false;

        }

        function UpdatePageList() {

            $('#form1').attr('action', '/anetn/Student/StUnitList/Update');
            $('#form1').removeAttr('target');
            $('#form1').submit();
            return false;
        }

    
