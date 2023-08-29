package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M39 187q0-43 14-74t35-50t46-32t47-17t35-7t15-1h189v4q0 8-6 13t-16 9t-21 4t-21 3l-17 2q-10 1-14 4q36 19 46 50t10 76q0 51-17 80t-39 50q-13 11-22 21t-9 23t12 25t29 27l28 26q15 14 28 33t21 40t8 49q0 113-96 160q-37 18-72 22t-51 4h-10q-8 0-25-1t-39-6t-43-15t-41-26t-31-43T.5 578.5T13 518t34-42t47-27t50-15t47-7t35-1h2q-12-15-17-30t-7-28t-1-25q-2 1-11 1q-15 0-41-4t-51-21q-61-40-61-132zm306 371q-2-41-35-68t-86-27h-13q-27 2-51 11t-42 27t-27 36t-8 40.5T96 617t31 28t44 16t54 4q56-5 90-34t30-73zm-50-411q-14-52-38-71t-62-18q-11 0-16 2q-34 9-48.5 50t-1.5 87q11 44 38 71t58 28q5 0 9-1t7-1q17-5 30-19t21-35t9-44t-6-49zm445 139v83H616v123h-82V369H410v-83h124V163h82v123h124z"/>`),
		g.Group(children),
	)
}