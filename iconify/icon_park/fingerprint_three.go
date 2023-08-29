package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42.4307 12.0391C37.702 7.38257 31.1539 4.5 23.9184 4.5C16.7254 4.5 10.2118 7.34876 5.48975 11.9571"/><path d="M6.6665 29.4745V29.4168C6.6665 19.8439 14.4269 12.0835 23.9998 12.0835"/><path d="M31.1694 13.6309C37.1649 16.3582 41.3333 22.4006 41.3333 29.4166V29.4296"/><path d="M14.25 36.9998V29.4165C14.25 24.0317 18.6152 19.6665 24 19.6665C29.3848 19.6665 33.75 24.0317 33.75 29.4165V36.9998"/><path d="M17.5259 43.5C19.4888 43.0154 20.7498 40.9456 20.7498 39.196C20.7498 37.3354 20.7498 34.4367 20.7498 30.5C20.7498 28.7051 22.2049 27.25 23.9998 27.25C25.7947 27.25 27.2498 28.7051 27.2498 30.5V39.196"/></g>`),
		g.Group(children),
	)
}