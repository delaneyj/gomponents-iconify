package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xmcc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm6.682-25h-5.16l-1.534 2.652L14.459 7H9.297L4 16.175l5.129 8.88l3.477-6.019L16 24.896l3.384-5.86l3.485 6.06L28 16.216L22.682 7zM9.122 21.544L6.02 16.178l2.467-4.272l3.1 5.37l-2.466 4.273v-.005zm.384-11.402l.818-1.421h3.091l1.555 2.693l-2.364 4.099l-3.1-5.371zm6.48 11.232l-2.362-4.102l2.364-4.097L17.044 15l1.31 2.273l-2.366 4.097l-.002.005zm1.017-9.96l1.558-2.693h3.091l.828 1.42l-3.12 5.372l-2.357-4.1zm7.968 6.537l-2.093 3.631l-2.488-4.32l3.1-5.359l.836 1.44l1.653 2.863l-1.008 1.745z"/>`),
		g.Group(children),
	)
}