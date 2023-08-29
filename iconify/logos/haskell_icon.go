package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HaskellIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#453A62" d="m0 180.591l60.235-90.294L0 0h45.176l60.236 90.297l-60.236 90.294z"/><path fill="#5E5086" d="m60.235 180.591l60.236-90.294L60.235 0h45.177l120.465 180.591h-45.171l-37.645-56.432l-37.652 56.432z"/><path fill="#8F4E8B" d="m205.804 127.92l-20.079-30.1H256v30.102h-50.196v-.002Zm-30.118-45.145l-20.078-30.1H256v30.1h-80.314Z"/>`),
		g.Group(children),
	)
}