package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lighttpd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#316682" d="m46.081 194.998l19.847-18.585l-.138-.87l-15.164-10.227l-4.643 27.742l.098 1.94"/><path fill="#8194A6" d="m46.081 194.998l5.015-28.526l1.176-.216L251.732 4.664L256 0l-7.561 4.193L31.524 151.445l-.177 1.234l14.734 42.319"/><path fill="#D9E2E9" d="M31.347 152.679L256 0L0 123.839l31.347 28.84"/><path fill="#D7E0E5" d="m117.239 210.672l-66.143-44.2L256 0L117.239 210.672"/>`),
		g.Group(children),
	)
}