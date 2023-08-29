package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Url(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M341 537q3 1 4 4t-1 6l-83 83q-20 20-46 32t-53 14t-54-4t-49-25q-27-21-42-51T1 536t9-61t34-54l129-129q31-31 72-46t84-8q38 6 60 22t35 37q2 5-.5 7.5T418 309t-7 6l-7 7q-10 11-24 11t-25.5-11.5T329 304t-31-5t-30 5t-27 18L93 470q-13 13-19 30t-4 34t10 34t26 27q25 17 54 12t49-24l64-63q2-3 6-1q30 13 62 18zM619 31q27 21 42 50t17 60t-9 62t-35 54L505 386q-31 31-72 45t-85 8q-38-5-59-22t-35-37q-2-4 1-7l18-18q11-10 25-10t24 10q12 12 27 18t31 6t31-6t26-18l148-147q13-13 19-31t4-35t-11-33t-26-27q-24-16-53-12t-50 25l-63 63q-3 3-6 1q-28-12-62-19q-4 0-4-4t1-5l83-83q20-20 46-32t52-15t54 5t50 25z"/>`),
		g.Group(children),
	)
}