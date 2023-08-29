package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BcaMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.77 20.331l-.009.045c-.352 1.997-2.256 3.616-4.254 3.616h0c-1.997 0-3.33-1.619-2.978-3.616l.65-3.685c.352-1.997 2.256-3.616 4.253-3.616h0c1.998 0 3.331 1.62 2.98 3.616l-.009.045m-13.586 1.798c1.507 0 2.514 1.222 2.248 2.73s-1.703 2.728-3.21 2.728H10.35l1.925-10.917h4.504c1.507 0 2.514 1.222 2.248 2.73s-1.703 2.729-3.21 2.729h0Zm0 0h-4.503m25.791 1.806h-4.729m-1.818 3.62l5.468-10.885l1.623 10.917M20.656 34.925h0c-1.017 0-1.696-.824-1.517-1.841l.211-1.197c.18-1.016 1.15-1.84 2.166-1.84h0c1.017 0 1.695.824 1.516 1.84l-.21 1.197c-.18 1.017-1.15 1.84-2.166 1.84Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M10.036 31.887c.179-1.016 1.148-1.84 2.165-1.84h0c1.017 0 1.696.824 1.516 1.84l-.535 3.038m-2.822-4.879l-.86 4.879"/><path d="M13.717 31.887c.18-1.016 1.15-1.84 2.166-1.84h0c1.017 0 1.696.824 1.516 1.84l-.535 3.038"/></g><ellipse cx="31.124" cy="27.791" fill="currentColor" rx=".806" ry=".676" transform="rotate(-42.481 31.124 27.791)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.727 30.046l-.861 4.879m3.234-7.364l-1.136 6.443c-.09.509.25.92.758.92h.276m4.86-.928c-.415.555-1.078.929-1.763.929h0c-1.017 0-1.696-.824-1.517-1.841l.211-1.197c.18-1.016 1.15-1.84 2.166-1.84h0c1.017 0 1.696.824 1.516 1.84l-.105.599h-3.682m-9.786-.599c.18-1.016 1.149-1.84 2.166-1.84h0c1.016 0 1.695.824 1.516 1.84l-.211 1.197c-.18 1.017-1.149 1.84-2.166 1.84h0c-1.016 0-1.695-.823-1.516-1.84m-.324 1.841l1.298-7.364"/>`),
		g.Group(children),
	)
}