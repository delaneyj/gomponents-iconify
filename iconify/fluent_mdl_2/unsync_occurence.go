package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsyncOccurence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1933 1843l-90 90l-235-234q-119 103-267 162t-317 59q-119 0-231-30t-212-87t-182-136t-143-181v306H128v-512h512v128H359q51 89 122 160t158 120t183 77t202 27q141 0 266-48t228-136L439 530q-63 73-106 160t-63 187l-126-25q23-119 76-222t129-190L115 205l90-90l1728 1728zM1024 256q-97 0-187 24t-171 67l-94-95q101-60 214-92t238-32q118 0 230 30t211 87t183 137t144 180V256h128v512h-512V640h281q-51-88-122-159t-158-121t-184-77t-201-27zm773 1221l-96-96q26-49 46-100t31-109l125 24q-14 77-41 146t-65 135z"/>`),
		g.Group(children),
	)
}