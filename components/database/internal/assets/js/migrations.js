$(document).ready(function () {
    hljs.initHighlightingOnLoad();

    $('#migrations tbody tr.migration button').click(function() {
        $(this).find('i')
            .toggleClass('glyphicon glyphicon-eye-open')
            .toggleClass('glyphicon glyphicon-eye-close');

        $(this).closest('#migrations tbody tr').next().toggle();
    });

    $('#migrations tbody button.close').click(function() {
        $(this).closest('#migrations tbody tr').prev().find('button').click();
    });
});