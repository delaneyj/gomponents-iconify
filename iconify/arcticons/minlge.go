package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minlge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.17 21.333V9.975l5.632 11.286l5.631-11.286v11.358"/><circle cx="26.922" cy="11.462" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.927 13.927v7.406m7.903 0v-4.656a2.816 2.816 0 1 0-5.631 0v4.656m-.001-4.656v-2.821M8.5 23.169v11.046h5.511m10.004-1.333a2.137 2.137 0 0 1-1.894 1.38a2.546 2.546 0 0 1-2.228-2.761v-1.795a2.546 2.546 0 0 1 2.228-2.762a2.546 2.546 0 0 1 2.228 2.762v.966h-4.456"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.756 28.412a2.122 2.122 0 0 1 1.88-1.381a2.54 2.54 0 0 1 2.211 2.761v1.795a2.54 2.54 0 0 1-2.211 2.762a2.007 2.007 0 0 1-2.028-2.269c0-1.69 1.375-1.407 4.285-1.407m19.194 2.5a2.666 2.666 0 0 1-2.343 1.381a2.767 2.767 0 0 1-2.756-2.762v-1.795a2.756 2.756 0 1 1 5.512 0v.967h-5.511m-2.044-3.985v8.284a2.767 2.767 0 0 1-2.755 2.762a2.31 2.31 0 0 1-1.929-.828"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.19 26.979a2.767 2.767 0 0 1 2.755 2.762v1.794a2.756 2.756 0 1 1-5.511 0v-1.794a2.767 2.767 0 0 1 2.755-2.762Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}