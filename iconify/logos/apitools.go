package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apitools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M177.616 77.951V63.768H78.327v14.183H64.143v106.373h28.368v-28.367h21.277v28.367h28.367v-28.367h21.276v28.367H191.8V77.95h-14.183zM92.51 92.134h70.92v35.458h-70.92V92.134zM14.509 0h226.922v14.163H14.509V0zm0 241.797h226.922V256H14.509v-14.203zM0 14.528h14.144v226.904H0V14.528zm241.796 0H256v226.904h-14.204V14.528z" fill="#2CB0AD"/>`),
		g.Group(children),
	)
}