package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 16.5a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0 0V24m-3.5-3.5h7m7.997-20l-7.722 7.721M23.497.5c-.412.411-1.171.676-1.85.84a8 8 0 0 1-2.775.153c-.72-.082-1.512-.243-1.832-.563M23.497.5c-.412.411-.676 1.17-.84 1.85a8 8 0 0 0-.153 2.775c.082.72.243 1.512.562 1.831M.502.5l7.722 7.722M.502.5c.411.411 1.17.676 1.85.84a8 8 0 0 0 2.775.153c.72-.082 1.512-.243 1.831-.563M.502.5c.411.411.675 1.17.84 1.85a8 8 0 0 1 .152 2.775c-.081.72-.243 1.512-.562 1.831M3.5 7.5l4-4"/>`),
		g.Group(children),
	)
}