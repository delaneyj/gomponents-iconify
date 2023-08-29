package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M902 702L577 833v-96l294-108q24-11 25.5-24.5t-21-21.5t-55-6t-57.5 13l-186 51v-96q74-33 224-32q40 0 58 .5t45.5 5T955 533q44 17 60.5 37t6 41.5t-40 43.5t-79.5 47zM705 449q-21 0-32-.5t-29-2t-27.5-6t-20-11t-15-18T577 385V160q0-13-9.5-22.5T545 128t-22.5 9.5T513 160v673l-128-64V0l128 49q24 14 64 31q85 37 138.5 99.5T769 321v32q0 14-.5 22.5t-3.5 24t-9.5 25t-19.5 17t-31 7.5zM71 717q-44-17-60.5-37t-6-41.5t40-44T124 547l197-79v92l-166 61q-30 19-25 35.5t31 16.5q24 0 64-7.5t68-14.5l28-7v89q-45 4-96 4q-40 0-58-1t-45.5-5.5T71 717z"/>`),
		g.Group(children),
	)
}