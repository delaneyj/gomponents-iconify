package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 389q-42 63-95 117t-115 100q1 14 1 27t1 28q0 126-27 249t-78 238q-74 167-185 298t-251 223t-307 139t-348 48q-172 0-335-47T0 1667q49 6 100 6q143 0 276-46t246-134q-67-1-129-22t-113-60t-90-92t-60-117q20 3 39 5t40 2q56 0 110-15q-74-15-135-53t-107-92t-70-123t-25-144v-5q88 50 191 53q-44-30-78-68t-59-84t-37-95t-13-103q0-56 14-109t43-102q80 99 178 177t208 135t232 88t248 39q-6-23-8-47t-3-49q0-87 33-163t90-134t133-90t164-33q88 0 167 34t140 99q71-14 137-39t129-63q-23 73-70 133t-114 99q126-15 241-66z"/>`),
		g.Group(children),
	)
}