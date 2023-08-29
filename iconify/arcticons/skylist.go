package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skylist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.016 5.821v18.916c3.058-3.446 8.362-7.458 16.544-12.009q.216-.334.45-.658l-.08.084l-2.583-1.28a12.41 12.41 0 0 1 5.397-2.105l-2.949-2.948Zm0 25.017V42.5h25.653V15.41q-.667.354-1.366.641l.12-.029c-1.258 1.268-1.816 2.443-4.828 3.945l.185-2.877l.122-.029q-.402.041-.806.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.56 12.73C4.428 24.482 2.425 32.66 9.016 36.702C4.599 29.05 19.443 21.525 28.096 17.12"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.326 6.01a15.015 15.015 0 0 0-16.766 6.72a14.906 14.906 0 0 0 1.06 2.316a14.907 14.907 0 0 0 1.476 2.075a15.214 15.214 0 0 0 7.176-1.995a15.182 15.182 0 0 0 7.054-9.117Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.423 16.021c-1.258 1.268-1.815 2.443-4.828 3.945l.185-2.876m.397-8.424c-1.728.456-3.024.351-5.83 2.209l2.582 1.278"/><circle cx="34.831" cy="10.306" r="1.614" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="rotate(-89.671 34.83 10.305)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.326 6.009a15.127 15.127 0 0 0-4.992-.45a7.687 7.687 0 0 0 .937 2.76a7.685 7.685 0 0 0 1.954 2.218a15.124 15.124 0 0 0 2.101-4.528ZM18.281 35.925h11.323m-11.323-5.149h11.323m-11.323-5.149h11.323"/>`),
		g.Group(children),
	)
}