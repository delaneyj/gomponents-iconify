package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RevolutJunior(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.585 10.34c-.21.94-2.087 12.946-2.087 12.946h2.822c3.656 0 6.372-3.023 6.372-7.62c0-3.233-1.88-5.325-4.798-5.325h-2.309Zm17.527 3.86c.01.162.015.324.017.486c.004 5.333-3.252 9.223-7.743 11.654h0l1.677 4.442l2.392 6.336c1.151 2.927 1.88 4.069 3.128 5.01a3.698 3.698 0 0 1-.94.416a18.25 18.25 0 0 1-6.266.94c-1.046 0-2.272-.13-3.458-.9s-2.333-2.177-3.22-4.732l-3.023-8.771h-3.022s-.954 6.272-1.353 8.348a18.025 18.025 0 0 0 0 4.595h0A19.466 19.466 0 0 1 11 43.376c-1.564 0-3.022-.522-3.128-3.233c.018-.734.087-1.465.209-2.188l3.857-24.833c.321-1.896.496-3.815.523-5.738c.004-.511.034-2.045.034-2.045c1.258-.313 5.503-1.059 14.274-.776c1.342.018 2.676.16 3.987.422"/><circle cx="36.274" cy="8.857" r="6.582" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.274 4.857v6a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-.664"/>`),
		g.Group(children),
	)
}