package main

// PrettifyPlaceholder 占位符
func PrettifyPlaceholder() string {
	return `<div style="height:25px;"> &nbsp; </div>`
}

// PrettifyQuote 引用
func PrettifyQuote(quote string) string {
	return `<section ><section xmlns="http://www.w3.org/1999/xhtml"  style="line-height: 27.2px;letter-spacing: 0.544px;"><section data-role="outer" label=""  style="font-size: 16px;"><section data-role="outer" label="" ><section class="" data-tools="" data-id="89124"  style="border-width: 0px;border-style: none;border-color: initial;"><section class=""  style="margin-right: auto;margin-left: auto;padding: 8px;text-align: left;min-height: 120px;background-image: url(&quot;https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5va3ugrzvPAfYDs7ZbyMFtlWmUvFTxeokxic0tkjDicoicRff6DP92M7XXjIFsCU7KOyBa8VwbDmIOMTA/640?wx_fmt=png&quot;);background-size: 100%;background-repeat: repeat-y;"><section class=""  style="padding: 5px 10px;line-height: 30px;font-size: 14px;"><p style="margin-top: 10px;margin-right: 8px;margin-left: 8px;text-align: justify;letter-spacing: 0.5px;line-height: 1.75em;text-indent: 2em;color: rgb(0, 0, 0);">` + quote + `<br></p></section></section></section></section></section><section data-role="outer" label="" ><section data-role="outer" label="" ><p><br ></p></section></section></section><section xmlns="http://www.w3.org/1999/xhtml"  style="line-height: 27.2px;letter-spacing: 0.544px;"><section data-role="outer" label=""  style="font-size: 16px;"><section data-role="outer" label="" ><p style="margin-right: 16px;margin-bottom: 15px;margin-left: 16px;text-align: center;"><img class="" data-copyright="0" data-ratio="0.02894736842105263" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZBVI2hK7lm0dqfVOovCbq3M0HfNwCnsKibw2mfQJ8unOSTr78wEvlibYwNbqGaIvRXYGtnAxRALzEA/640?wx_fmt=png" data-type="png" data-w="1140" style="letter-spacing: 0.544px; visibility: visible !important; width: 677px !important; height: auto !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZBVI2hK7lm0dqfVOovCbq3M0HfNwCnsKibw2mfQJ8unOSTr78wEvlibYwNbqGaIvRXYGtnAxRALzEA/640?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"></p></section></section></section></section>`
}

// PrettifySubtitle 标题
func PrettifySubtitle(title string) string {
	return `<p style="margin-right:16px; margin-left:16px; margin-bottom:25px; white-space:normal; letter-spacing:0.5px; line-height:1.75em;"><span style="clear:both; min-height:1em; line-height:1.75em; color:#485BF7; font-size:16px;"><strong>` + title + `</strong></span></p>`
}

// PrettifyContent 内容
func PrettifyContent(content string) string {
	return `<p style="margin-right:16px; margin-left:16px; margin-bottom:25px; white-space:normal; letter-spacing:0.5px; line-height:1.75em;"><span style="font-size:15px;">` + content + `</span></p>`
}

// PrettifyEnd 结尾
func PrettifyEnd() string {
	return `<section style="margin-right: 16px;margin-left: 16px;white-space: normal;"><img class="" data-ratio="0.06748466257668712" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZBVI2hK7lm0dqfVOovCbq3JIqoQylkUB8XBLWbmlvJkBIgZjOdsxVpGJibgibk7IJwnsHdibWw1ibAPw/640?wx_fmt=png" data-type="png" data-w="1141" style="text-align: center; font-family: mp-quote, -apple-system-font, BlinkMacSystemFont, &quot;Helvetica Neue&quot;, &quot;PingFang SC&quot;, &quot;Hiragino Sans GB&quot;, &quot;Microsoft YaHei UI&quot;, &quot;Microsoft YaHei&quot;, Arial, sans-serif; visibility: visible !important; width: 677px !important; height: auto !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZBVI2hK7lm0dqfVOovCbq3JIqoQylkUB8XBLWbmlvJkBIgZjOdsxVpGJibgibk7IJwnsHdibWw1ibAPw/640?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"></section>`
}

// PrettifyCopyright 版权
func PrettifyCopyright(content string) string {
	return `<p style="margin-bottom:25px; white-space:normal; letter-spacing:0.5px; line-height:1.75em; text-align:center;"><span style="font-size:15px;">` + content + `</span></p>`
}

// PrettifyFooter 页尾
func PrettifyFooter() string {
	return `<p style="text-align: center;"><img class="" data-ratio="0.15350877192982457" data-s="300,640" data-src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZVBLodFcP5fgucCKuEP0dCwxpQ6gHnKypJia9nSRjCLMA0MkrPM4VcKr9D6mGZEXBPCb2v1Bw7qLw/640?wx_fmt=png" data-type="png" data-w="1140" style="width: 677px !important; height: auto !important; visibility: visible !important;" _width="677px" src="https://mmbiz.qpic.cn/mmbiz_png/QicyPhNHD5vZVBLodFcP5fgucCKuEP0dCwxpQ6gHnKypJia9nSRjCLMA0MkrPM4VcKr9D6mGZEXBPCb2v1Bw7qLw/640?wx_fmt=png&amp;tp=webp&amp;wxfrom=5&amp;wx_lazy=1&amp;wx_co=1" crossorigin="anonymous" data-fail="0"><br></p>`
}
