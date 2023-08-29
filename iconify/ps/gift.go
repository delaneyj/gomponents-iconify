package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 128h-76q12-20 12-43q0-35-25-60T320 0q-40 0-64 30q-24-30-64-30q-35 0-60 25t-25 60q0 23 12 43H43q-18 0-30.5 12.5T0 171v42q0 18 12.5 30.5T43 256v192q0 27 18 45.5t46 18.5h298q28 0 46-18.5t18-45.5V256q18 0 30.5-12.5T512 213v-42q0-18-12.5-30.5T469 128zM192 469h-85q-22 0-22-21V256h107v213zm0-256H43v-42h149v42zm0-85q-17 0-30-12.5T149 85q0-17 13-29.5T192 43t30 12.5T235 85q0 18-13 30.5T192 128zm85 341h-42V256h42v213zm0-256h-42v-42h42v42zm0-128q0-17 13-29.5T320 43t30 12.5T363 85q0 18-13 30.5T320 128t-30-12.5T277 85zm150 363q0 21-22 21h-85V256h107v192zm42-235H320v-42h149v42z"/>`),
		g.Group(children),
	)
}