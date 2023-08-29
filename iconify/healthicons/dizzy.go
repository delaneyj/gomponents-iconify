package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M27.293 17.293a1 1 0 0 1 1.414 0l1.793 1.793l1.793-1.793a1 1 0 0 1 1.414 1.414L31.914 20.5l1.793 1.793a1 1 0 0 1-1.414 1.414L30.5 21.914l-1.793 1.793a1 1 0 0 1-1.414-1.414l1.793-1.793l-1.793-1.793a1.002 1.002 0 0 1 0-1.414Zm-13 0a1 1 0 0 1 1.414 0l1.793 1.793l1.793-1.793a1 1 0 0 1 1.414 1.414L18.914 20.5l1.793 1.793a1.002 1.002 0 0 1 0 1.414a1 1 0 0 1-1.414 0L17.5 21.914l-1.792 1.792v.001a1 1 0 0 1-1.415-1.414l1.793-1.793l-1.793-1.793a1.002 1.002 0 0 1 0-1.414ZM31 32c0 3.314-3.134 4-7 4s-7-.686-7-4s3.134-6 7-6s7 2.686 7 6Z"/><path fill-rule="evenodd" d="M24 42c9.941 0 18-8.059 18-18S33.941 6 24 6S6 14.059 6 24s8.059 18 18 18Zm0-2c8.837 0 16-7.163 16-16S32.837 8 24 8S8 15.163 8 24s7.163 16 16 16Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}