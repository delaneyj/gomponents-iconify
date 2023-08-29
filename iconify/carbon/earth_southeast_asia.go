package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthSoutheastAsia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 24h3v2h-3zm3-2h2v2h-2zm0-4h2v2h-2z"/><path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14.016 14.016 0 0 0 16 2Zm-.603 2.03l2.507 5.016L16.323 13H10.27l-1.41 4.387l2.852 2.852l-1.007 1.512L6.286 19H4.395A11.936 11.936 0 0 1 15.397 4.03ZM5.106 21h.608l3.935 2.449a1.991 1.991 0 0 0 2.721-.589l1.008-1.512a2.005 2.005 0 0 0-.25-2.523l-1.989-1.989l.59-1.836h4.594a1.99 1.99 0 0 0 1.857-1.257l1.581-3.953a2.002 2.002 0 0 0-.067-1.638l-2.01-4.02A11.96 11.96 0 0 1 26.894 21H23a2.002 2.002 0 0 0-2 2v3.894A11.948 11.948 0 0 1 5.105 21ZM23 25.727V23h2.727A12.076 12.076 0 0 1 23 25.727Z"/>`),
		g.Group(children),
	)
}