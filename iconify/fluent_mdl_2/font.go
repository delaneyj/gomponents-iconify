package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M765 1024H387l-85 256H167L509 256h134l268 802l-81 162l-65-196zm-43-128L576 458L430 896h292zm982 679q17 41 31 73t35 54t50 34t78 12v44h-449v-44h37q20 0 39-3t30-17t12-37q0-14-6-39t-16-54t-21-59t-22-58t-19-46t-11-27h-448q-3 8-12 27t-20 46t-24 57t-23 58t-17 54t-7 40q0 24 12 35t31 16t38 5t37 2v44H662v-44q49-9 76-21t44-32t30-52t33-79l392-924h82l385 935zm-291-295l-177-381l-169 381h346z"/>`),
		g.Group(children),
	)
}