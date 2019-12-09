package main

// PrettifyPlaceholder 占位符
func PrettifyPlaceholder() string {
	return `<div style="height:25px;">&nbsp;</div>`
}

// PrettifyQuote 引用
func PrettifyQuote(quote string) string {
	return `<section ><section xmlns="http://www.w3.org/1999/xhtml"  style="line-height: 27.2px;letter-spacing: 0.544px;"><section data-role="outer" label=""  style="font-size: 16px;"><section data-role="outer" label="" ><section class="" data-tools="" data-id="89124"  style="border-width: 0px;border-style: none;border-color: initial;"><section class="" style="margin-right: auto;margin-left: auto;padding: 8px;text-align: left;min-height: 120px;background-image: url(&quot;https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml9IzsrY6RowVxU28IdU4fvBS19K0v5ppt73nXzhEgoKwiaSUYB98ab8OA/0?wx_fmt=png&quot;);background-size: 100%;background-repeat: repeat-y;"><section class=""  style="padding: 5px 10px;line-height: 30px;font-size: 14px;"><p style="margin-top: 10px;margin-right: 8px;margin-left: 8px;text-align: justify;letter-spacing: 0.5px;line-height: 1.75em;text-indent: 2em;color: rgb(0, 0, 0);">` + quote + `<br></p></section></section></section></section></section><section data-role="outer" label="" ><section data-role="outer" label="" ><p><br ></p></section></section></section><section xmlns="http://www.w3.org/1999/xhtml"  style="line-height: 27.2px;letter-spacing: 0.544px;"><section data-role="outer" label=""  style="font-size: 16px;"><section data-role="outer" label="" ><p style="margin-right: 16px;margin-bottom: 15px;margin-left: 16px;text-align: center;"><img class="" data-copyright="0" data-ratio="0.02894736842105263" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml91QDZxIYZnNVvsibkGmhIv29Ea4r7v8jicxc7axGw7icMibBrjfbZhn2Q8Q/0?wx_fmt=png" data-type="png" data-w="1140" style="letter-spacing: 0.544px; visibility: visible !important; width: 677px !important; height: auto !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml91QDZxIYZnNVvsibkGmhIv29Ea4r7v8jicxc7axGw7icMibBrjfbZhn2Q8Q/0?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"></p></section></section></section></section>`
}

// PrettifyAuthor 作者
func PrettifyAuthor(author string) string {
	return `<section mpa-from-tpl="t" style="margin-right: 16px;margin-bottom: 5px;margin-left: 16px;font-size: 16px;text-align: right;letter-spacing: 0.5px;line-height: 1.75em;"><strong mpa-from-tpl="t" style="color: rgb(102, 102, 102);font-size: 13px;"><strong mpa-from-tpl="t">文&nbsp;|</strong></strong>&nbsp;<span style="line-height: 1.75em;font-size: 13px;color: rgb(102, 102, 102);">` + author + `</span></section>`
}

// PrettifyCopyright 版权
func PrettifyCopyright(copyright string) string {
	return `<section mpa-from-tpl="t" style="margin-right: 16px;margin-bottom: 5px;margin-left: 16px;font-size: 16px;text-align: right;letter-spacing: 0.5px;line-height: 1.75em;"><strong mpa-from-tpl="t" style="color: rgb(102, 102, 102);font-size: 13px;"><strong mpa-from-tpl="t">排版&nbsp;|</strong></strong>&nbsp;<span style="line-height: 1.75em;font-size: 13px;color: rgb(102, 102, 102);">` + copyright + `</span></section>`
}

// PrettifySubtitle 标题
func PrettifySubtitle(title string) string {
	return `<p style="margin-right:16px; margin-left:16px; margin-bottom:25px; white-space:normal; letter-spacing:0.5px; line-height:1.75em;"><span style="clear:both; min-height:1em; line-height:1.75em; color:rgb(72, 91, 247); font-size:16px;"><strong>` + title + `</strong></span></p>`
}

// PrettifyContent 内容
func PrettifyContent(content string) string {
	return `<p style="margin-right:16px; margin-left:16px; margin-bottom:25px; white-space:normal; letter-spacing:0.5px; line-height:1.75em;"><span style="font-size:15px;">` + content + `</span></p>`
}

// PrettifyEnd 结尾
func PrettifyEnd() string {
	return `<section style="margin-right: 16px;margin-left: 16px;white-space: normal;"><img class="" data-ratio="0.06748466257668712" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml970FaCpxFswmsq5iaWPiaAzb789Mn4cJMock5SKPmDyquP3SsLj1dkx5A/0?wx_fmt=png" data-type="png" data-w="1141" style="text-align: center; font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif; visibility: visible !important; width: 677px !important; height: auto !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml970FaCpxFswmsq5iaWPiaAzb789Mn4cJMock5SKPmDyquP3SsLj1dkx5A/0?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"></section>`
}

// PrettifyMore 更多
func PrettifyMore() string {
	return `<p><img class="" data-ratio="0.1270815074496056" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml9aicPTSFC78q8HKZsVhsqXF7giaZaMXX4h8shzCGDy4oSqmZEgrdFbKLg/0?wx_fmt=png" data-type="png" data-w="1141" style="visibility: visible !important; width: 677px !important; height: auto !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml9aicPTSFC78q8HKZsVhsqXF7giaZaMXX4h8shzCGDy4oSqmZEgrdFbKLg/0?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"></p>`
}

// PrettifyStrong 重点
func PrettifyStrong(strong string) string {
	return `<p style="margin-right:16px; margin-left:16px; margin-bottom:0px; white-space:normal; letter-spacing:0.5px; line-height:1.75em;"><span style="clear:both; min-height:1em; line-height:1.75em; color:rgb(72, 91, 247); font-size:16px;"><strong>` + strong + `</strong></span></p>`
}

// PrettifyFooter 页尾
func PrettifyFooter() string {
	return `<p style="text-align: center;"><img class="" data-ratio="0.15350877192982457" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml9nticB5HwkCqNUkezxvhwZ3pDOSviaQ7icB0YgOJOIYQaPQHMH2ibdSpEIg/0?wx_fmt=png" data-type="png" data-w="1140" style="width: 677px !important; height: auto !important; visibility: visible !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/TnrI8Zz76vdsJXwDg7F7pNicFLVmTKml9nticB5HwkCqNUkezxvhwZ3pDOSviaQ7icB0YgOJOIYQaPQHMH2ibdSpEIg/0?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"><br></p>`
}
