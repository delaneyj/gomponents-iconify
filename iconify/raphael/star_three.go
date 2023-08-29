package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.44 28.18c-.418 0-.834-.13-1.188-.39l-5.75-4.248L9.75 27.79a2.003 2.003 0 0 1-3.087-2.243l2.26-6.783l-5.815-4.158a2.001 2.001 0 0 1 1.164-3.628h.014l7.15.056l2.157-6.816a1.999 1.999 0 0 1 3.811.001l2.155 6.816l7.15-.056h.015c.867 0 1.636.556 1.903 1.382c.27.83-.028 1.737-.74 2.246l-5.814 4.158l2.263 6.783a2.003 2.003 0 0 1-1.896 2.634z"/>`),
		g.Group(children),
	)
}