package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Placealtadd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M764 765q-14 21-28 35l-289 224l-287-224q-14-14-28-35q-62-62-97-144T0 448t35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448t-35 173t-97 144zM448 128q-87 0-160.5 43t-116 116.5T129 448t42.5 160.5t116 116T448 767t160.5-42.5t116.5-116T768 448t-43-160.5T608.5 171T448 128zm192 384H512v128q0 26-19 45t-45.5 19t-45-19t-18.5-45V512H256q-27 0-45.5-19T192 447.5t18.5-45T256 384h128V256q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v128h128q26 0 45 18.5t19 45t-19 45.5t-45 19z"/>`),
		g.Group(children),
	)
}