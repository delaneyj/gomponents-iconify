package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntimatterDimensions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.422" cy="12.73" r="4.683" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.91 8.193a4.684 4.684 0 0 0-8.347 1.067m6.498 5.35a4.688 4.688 0 0 0-5.029-1.082m-2.922 3.567a4.692 4.692 0 0 0 1.7 4.466"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.226 16.45a4.732 4.732 0 0 0 .045-.651a4.683 4.683 0 1 0-7.852 3.447"/><circle cx="26.243" cy="23.56" r="4.683" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.838 20.098a4.684 4.684 0 0 0-8.84-.88"/><circle cx="9.071" cy="37.344" r="4.571" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.359 16.933a4.684 4.684 0 0 0 2.063 8.888a4.742 4.742 0 0 0 .645-.044"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.753 22.228a4.683 4.683 0 1 0-4.433 8.053m5.939-3.112a4.7 4.7 0 0 0-1.893 2.185m.91 5.039a4.683 4.683 0 1 0 6.389-6.825"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.197 32.418a4.682 4.682 0 1 0 .664-8.093"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.529 28.845a4.684 4.684 0 1 0-3.386-4.783m1.217-6.436a4.684 4.684 0 1 0-5.887-6.235"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.791 9.616a3.225 3.225 0 1 0-5.877 2.176M10.937 25.105a3.224 3.224 0 0 0 4.8 4.3"/><circle cx="19.87" cy="32.249" r="3.224" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.838 35.333a3.225 3.225 0 1 0 6.059-2.187"/>`),
		g.Group(children),
	)
}