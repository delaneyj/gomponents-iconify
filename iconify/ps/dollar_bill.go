package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DollarBill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21zM43 64h42q0 17-12.5 30T43 107V64zm0 256v-43q17 0 29.5 13T85 320H43zm234-9v-76q0-22-21-22t-21 22v76q-22-17-22-55q0-26 13-45t30-19t30 19t13 45q0 35-22 55zm192 9h-42q0-17 12.5-30t29.5-13v43zm0-85q-35 0-60 25t-25 60h-60q17-26 17-64q0-45-25-76t-60-31t-60 31t-25 76q0 38 17 64h-60q0-35-25-60t-60-25v-86q35 0 60-25t25-60h256q0 35 25 60t60 25v86zm0-128q-17 0-29.5-13T427 64h42v43zM341 85H171q-22 0-22 22q0 9 6 15t16 6h170q10 0 16-6t6-15q0-22-22-22z"/>`),
		g.Group(children),
	)
}