package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1504 1600q14 0 23 9t9 23v128q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-128q0-14 9-23t23-9h1472zm-1374-64q3-55 16-107t30-95t46-87t53.5-76t64.5-69.5t66-60t70.5-55T543 939t65-43q-43-28-65-43t-66.5-47.5t-70.5-55t-66-60t-64.5-69.5t-53.5-76t-46-87t-30-95t-16-107h1276q-3 55-16 107t-30 95t-46 87t-53.5 76t-64.5 69.5t-66 60t-70.5 55T993 853t-65 43q43 28 65 43t66.5 47.5t70.5 55t66 60t64.5 69.5t53.5 76t46 87t30 95t16 107H130zM1504 0q14 0 23 9t9 23v128q0 14-9 23t-23 9H32q-14 0-23-9t-9-23V32Q0 18 9 9t23-9h1472z"/>`),
		g.Group(children),
	)
}