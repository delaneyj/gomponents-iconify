package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkinalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M764 765q-14 20-28 35l-289 224l-287-224q-14-15-28-35q-62-62-97-144T0 448t35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95.5 143T896 448t-35 173t-97 144zM448 128q-87 0-160.5 43T171 287.5T128 448t43 160.5t116.5 116T448 767t160.5-42.5t116.5-116T768 448t-43-160.5T608.5 171T448 128zm-.5 512Q368 640 312 583.5t-56-136T312 312t135.5-56t136 56T640 447.5t-56.5 136t-136 56.5z"/>`),
		g.Group(children),
	)
}