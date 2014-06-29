
/**
 * 修改管理员信息
 */

/**
 * 提交检测
 */
function form_submit() {
	var realname = $.trim($("#realname").val());
	if (realname == '') {
		$("#realname").focus();
		notice_tips("请输入真实姓名!");
		return false;
	}

//    var eamilPartten = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;
	var eamilPartten = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/;
	var email = $.trim($("#email").val());
	if (email == '') {
		$("#email").focus();
		notice_tips("请输入电子邮件!");
		return false;
	}
	if (!eamilPartten.test(email)) {
	    $("#email").focus();
        notice_tips("请输入有效电子邮件!");
        return false;
	}

	var lang = $.trim($("#lang").val());
	if (lang == '') {
		$("#lang").focus();
		notice_tips("请选择语言!");
		return false;
	}

	return true;
}