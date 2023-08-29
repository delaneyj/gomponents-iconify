package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v128H512V384h1536zM512 896h1536v128H512V896zm0 512h1536v128H512v-128zM135 349q-14 11-37 21t-43 11v-69q31-11 59-24t54-33h49v385h-82V349zm-1 420q25 0 47 6t39 20t26 34t10 47q0 31-12 54t-29 42t-39 33t-38 27t-30 25t-12 26h160v69H5v-41q0-18 10-39q15-34 41-57t51-44t45-40t19-47q0-27-15-38t-40-11q-26 0-49 11t-44 28v-73q50-32 111-32zm36 699q36 5 61 27t25 61q0 31-12 53t-33 35t-46 20t-54 7q-24 0-48-4t-47-15v-73q19 14 41 21t47 7q26 0 46-11t21-41q0-22-13-32t-31-15t-39-5t-34 0v-64h32q18 0 34-3t28-14t12-31q0-26-16-36t-40-10q-39 0-74 24v-68q22-11 45-15t48-5q22 0 44 5t39 16t28 30t11 43q0 38-19 60t-56 32v1z"/>`),
		g.Group(children),
	)
}