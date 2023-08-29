package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeywordSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.965C0 1.88.88 1 1.965 1h2.807c1.085 0 1.965.88 1.965 1.965v.561c0 1.086-.88 1.965-1.965 1.965H1.965A1.965 1.965 0 0 1 0 3.526v-.561Zm1.965-.28a.28.28 0 0 0-.28.28v.561a.28.28 0 0 0 .28.281h2.807a.28.28 0 0 0 .28-.28v-.562a.28.28 0 0 0-.28-.28H1.965Zm6.175.561c0-.465.377-.842.842-.842h6.176a.842.842 0 1 1 0 1.684H8.982a.842.842 0 0 1-.842-.842ZM.28 8.298c0-.465.378-.842.843-.842H11.79a.842.842 0 1 1 0 1.684H1.123a.842.842 0 0 1-.842-.842Zm0 5.052c0-.464.378-.841.843-.841h13.474a.842.842 0 1 1 0 1.684H1.123a.842.842 0 0 1-.842-.842Z" clip-rule="evenodd"/><path fill="currentColor" d="M14.877 9.14a.842.842 0 1 0 0-1.684a.842.842 0 0 0 0 1.684Z"/>`),
		g.Group(children),
	)
}