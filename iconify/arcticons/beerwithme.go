package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beerwithme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.957 28.231a7.064 7.064 0 0 1-13.913 0m-.001 0a22.256 22.256 0 0 1-.386-2.952a28.947 28.947 0 0 1 .125-4.7c.221-2.207.62-4.39 1.022-6.571q.77-4.168 1.566-8.33m11.587 22.553a22.256 22.256 0 0 0 .386-2.952a28.947 28.947 0 0 0-.125-4.7c-.221-2.207-.62-4.39-1.022-6.571q-.77-4.168-1.566-8.33m-9.26-.001c0-.65 2.073-1.177 4.63-1.177s4.63.527 4.63 1.177"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.273 11.475c0-.434 2.563-.785 5.724-.785s5.725.351 5.725.785M26.34 33.789a11.296 11.296 0 0 0-.585 1.332a8.742 8.742 0 0 0-.409 1.5a8.89 8.89 0 0 0-.109 1.788a4.451 4.451 0 0 0 .454 2.033a4.047 4.047 0 0 0 1.225 1.28a7.777 7.777 0 0 0 4.072 1.43m-9.328-9.363a11.296 11.296 0 0 1 .585 1.332a8.742 8.742 0 0 1 .409 1.5a8.89 8.89 0 0 1 .109 1.788a4.451 4.451 0 0 1-.453 2.033a4.047 4.047 0 0 1-1.226 1.28a7.777 7.777 0 0 1-4.072 1.43"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.012 43.151a70.266 70.266 0 0 0 7.11.349q3.44-.006 6.867-.349"/>`),
		g.Group(children),
	)
}