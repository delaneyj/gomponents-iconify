package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upperrightwhitecircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#BCBCBC" d="M478.609 222.542c0-104.009-84.316-188.325-188.325-188.325c-54.061 0-102.793 22.786-137.139 59.269l-.068-.068l-62.847 62.847l17.989 17.989a188.575 188.575 0 0 0-6.26 48.288c0 104.009 84.316 188.325 188.325 188.325c16.694 0 32.876-2.183 48.288-6.26l17.989 17.989l62.847-62.847l-.068-.068c36.483-34.346 59.269-83.079 59.269-137.139z"/><circle cx="223.395" cy="289.43" r="188.325" fill="#E0E0E0"/>`),
		g.Group(children),
	)
}