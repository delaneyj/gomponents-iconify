package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c.896 0 1.775.193 2.546.557c.348.165.669.362.955.587c.347.272.645.585.881.93c.432.627.644 1.336.616 2.052c-.028.716-.296 1.412-.776 2.017c-.48.605-1.154 1.096-1.952 1.42a6.075 6.075 0 0 1-2.583.43a5.865 5.865 0 0 1-2.497-.685c-.74-.402-1.332-.957-1.713-1.605M12 12H4m8 0h8m-3.476-5.703c-.381-.648-.973-1.202-1.714-1.605a5.866 5.866 0 0 0-2.496-.684a6.075 6.075 0 0 0-2.584.428c-.798.325-1.472.816-1.952 1.42c-.48.606-.747 1.303-.776 2.019c-.008.21.005.418.037.625"/>`),
		g.Group(children),
	)
}