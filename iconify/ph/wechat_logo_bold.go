package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WechatLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236.06 187.69a84 84 0 0 0-63.73-118.78a84 84 0 1 0-152.39 70.78l-7.24 25.36A18 18 0 0 0 30 188a17.8 17.8 0 0 0 5-.7l25.35-7.24a82.4 82.4 0 0 0 23.42 7a84 84 0 0 0 112 41l25.36 7.24a18 18 0 0 0 22.25-22.25Zm-168.78-31a12 12 0 0 0-9.05-1l-19.49 5.57l5.57-19.49a12 12 0 0 0-1-9.05A60 60 0 0 1 145 69.37a84 84 0 0 0-68.55 91.39a60.34 60.34 0 0 1-9.17-4.06Zm144.41 33.07l5.57 19.49l-19.49-5.57a12 12 0 0 0-9.05 1a60 60 0 1 1 24-24a12 12 0 0 0-1.03 9.09ZM152 144a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm48 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}