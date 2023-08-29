package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m92.038 24.333l-8.62-8.622a1.615 1.615 0 0 0-2.282 0L49.987 46.86l-6.07-6.068a1.614 1.614 0 0 0-2.282 0l-8.622 8.622a1.611 1.611 0 0 0 0 2.282l15.782 15.778c.302.302.712.473 1.141.473c.019 0 .037-.01.056-.01c.016 0 .033.009.05.009a1.61 1.61 0 0 0 1.141-.473l40.855-40.857c.63-.632.63-1.653 0-2.283z"/><path fill="currentColor" d="M72.022 53.625v21.159H27.978V30.74h31.06l9.979-9.978H23.193v.007c-.023 0-.044-.007-.068-.007a5.118 5.118 0 0 0-5.113 5H18v54h.013A5.111 5.111 0 0 0 23 84.749v.013h54v-.013a5.11 5.11 0 0 0 4.987-4.987H82V43.647l-9.978 9.978z"/>`),
		g.Group(children),
	)
}