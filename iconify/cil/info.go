package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 95.998h34.924v34.924H256z"/><path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M285.313 359.032a18.123 18.123 0 0 1-15.6 8.966a18.061 18.061 0 0 1-17.327-23.157l35.67-121.277a49.577 49.577 0 0 0-93.356-32.992l-11.718 28.234l29.557 12.266l11.718-28.235a17.577 17.577 0 0 1 33.1 11.7l-35.67 121.277A50.061 50.061 0 0 0 269.709 400a50.227 50.227 0 0 0 43.25-24.853l15.1-25.913l-27.646-16.115Z"/>`),
		g.Group(children),
	)
}