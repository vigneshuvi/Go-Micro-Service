/**
 * Created by Vignesh on 12/10/15.
 */
var uviWSUtil = {
    config: "",

    init: function () {

        uviWSUtil.getNWSMConfig();

        $('#uviCreate').on('click', function (e) {
            uviWSUtil.sampleCreateWS();
        });
        $('#uviUpdate').on('click', function (e) {
            uviWSUtil.sampleUpdateWS();
        });
        $('#uviShow').on('click', function (e) {
            uviWSUtil.sampleShowWS();
        });
        $('#uviShowAll').on('click', function (e) {
            uviWSUtil.sampleShowAllWS();
        });
        $('#uviDelete').on('click', function (e) {
            uviWSUtil.sampleDeleteWS();
        });
    },

    getNWSMConfig : function() {
           $.ajax(
            {
                method: "GET",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/u1v2i3C4o5n6f7i8g9",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    uviWSUtil.config= resData.environment;
                    $("#uviResult").html(resData.environment);
                }, 
                error: function(xhr, status, error) {
                    var err = eval("(" + xhr.responseText + ")");
                    console.log("API Error", err.Message);
                }
        });
    },

    getContentType :function () {
        return 'application/json';
    },

    getCurrentDate :function () {
        return new Date().getTime();
    },

    getRequestHeader : function () {

        if(uviWSUtil.config === "QA" && uviWSUtil.config === "PROD") {
                    console.log(uviWSUtil.environment);
            return { 'xauthorization': 'bypass', 'xdate' : uviWSUtil.getCurrentDate(), 'Access-Control-Allow-Origin':'*' };
        } else {

            return { 'xdate' : uviWSUtil.getCurrentDate(), 'Access-Control-Allow-Origin':'*' };
        }
    },

    sampleCreateWS : function () {
        $.ajax(
            {
                method: "POST",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/user",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    $("#uviResult").html(JSON.stringify(resData));
                }, 
                error: function(xhr, status, error) {
                   $("#uviResult").html(error);
                }
        });
    },
    sampleUpdateWS :function() {
         $.ajax(
            {
                method: "PUT",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/user?id=test",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    $("#uviResult").html(JSON.stringify(resData));
                }, 
                error: function(xhr, status, error) {
                    $("#uviResult").html(error);
                }
        });

    },
    sampleShowWS: function () {
         $.ajax(
            {
                method: "GET",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/user?id=test",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    $("#uviResult").html(JSON.stringify(resData));
                }, 
                error: function(xhr, status, error) {
                    $("#uviResult").html(error);
                }
        });

    },
    sampleShowAllWS: function () {
         $.ajax(
            {
                method: "GET",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/user",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    $("#uviResult").html(JSON.stringify(resData));
                }, 
                error: function(xhr, status, error) {
                    $("#uviResult").html(error);
                }
        });

    },
    sampleDeleteWS: function () {
         $.ajax(
            {
                method: "DELETE",
                headers: uviWSUtil.getRequestHeader(),
                url: "http://localhost:3000/v1/user?id=test",
                contentType: uviWSUtil.getContentType(),
                success: function(resData) {
                    $("#uviResult").html(JSON.stringify(resData));
                }, 
                error: function(xhr, status, error) {
                    $("#uviResult").html(error);
                }
        });
    }
};