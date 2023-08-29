package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FuneralUrn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f585id0)" d="M16 2a7 7 0 0 0-7 7h14a7 7 0 0 0-7-7Z"/><path fill="url(#f585id1)" d="M4 14c0-7 4-7 4-7h16s4 0 4 7c0 6.538-3.167 12.021-4.899 14.531c-.663.961-1.773 1.469-2.94 1.469H11.84c-1.168 0-2.278-.508-2.941-1.469C7.167 26.021 4 20.538 4 14Z"/><path fill="url(#f585id2)" fill-rule="evenodd" d="M27.97 13a11.926 11.926 0 0 0-.278-2H4.308a11.91 11.91 0 0 0-.279 2h23.942ZM4.026 15c.033.68.099 1.348.192 2h23.566c.093-.652.16-1.32.192-2H4.026Z" clip-rule="evenodd"/><defs><radialGradient id="f585id0" cx="0" cy="0" r="1" gradientTransform="matrix(-3 6 -12 -5.99999 18 3.5)" gradientUnits="userSpaceOnUse"><stop stop-color="#F3AD61"/><stop offset="1" stop-color="#A56953"/></radialGradient><radialGradient id="f585id1" cx="0" cy="0" r="1" gradientTransform="matrix(-5.92854 16.57143 -20.23768 -7.24016 19.429 11.929)" gradientUnits="userSpaceOnUse"><stop stop-color="#F3AD61"/><stop offset="1" stop-color="#A56953"/></radialGradient><linearGradient id="f585id2" x1="3.5" x2="22.5" y1="12" y2="12" gradientUnits="userSpaceOnUse"><stop stop-color="#7D4533"/><stop offset="1" stop-color="#A56953"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}