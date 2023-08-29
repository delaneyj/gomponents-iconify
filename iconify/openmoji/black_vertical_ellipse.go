package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackVerticalEllipse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M56 35.959C56.035 51.422 47.109 63.977 36.063 64C25.017 64.023 16.034 51.506 16 36.041C15.965 20.578 24.891 8.023 35.937 8C46.983 7.977 55.966 20.494 56 35.959Z"/><path fill="#3F3F3F" d="M56 35.959C56.035 51.422 47.109 63.977 36.063 64C25.017 64.023 16.034 51.506 16 36.041C15.965 20.578 24.891 8.023 35.937 8C46.983 7.977 55.966 20.494 56 35.959Z"/><path fill="none" stroke="#000" stroke-width="2" d="M56 36c0 15.464-8.954 28-20 28S16 51.464 16 36S24.954 8 36 8s20 12.536 20 28Z"/>`),
		g.Group(children),
	)
}