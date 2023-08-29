package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JiraWordmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><linearGradient id="deviconJiraWordmark0" x1="22.034" x2="17.118" y1="9.773" y2="14.842" gradientTransform="translate(1.136 -13.247) scale(3.90869)" gradientUnits="userSpaceOnUse"><stop offset=".176" stop-color="#0052cc"/><stop offset="1" stop-color="#2684ff"/></linearGradient><linearGradient id="deviconJiraWordmark1" x1="16.641" x2="10.957" y1="15.564" y2="21.094" gradientTransform="translate(1.136 -13.247) scale(3.90869)" gradientUnits="userSpaceOnUse"><stop offset=".176" stop-color="#0052cc"/><stop offset="1" stop-color="#2684ff"/></linearGradient></defs><path fill="#2684ff" d="M106.691 2.387h-45.16c0 11.258 9.114 20.37 20.367 20.37h8.309v8.04c0 11.258 9.113 20.371 20.371 20.371V6.273a3.89 3.89 0 0 0-3.887-3.886zm0 0"/><path fill="url(#deviconJiraWordmark0)" d="M84.313 24.902h-45.16c0 11.258 9.109 20.368 20.367 20.368h8.308v8.042c0 11.258 9.113 20.372 20.371 20.372V28.789a3.89 3.89 0 0 0-3.886-3.887zm0 0"/><path fill="url(#deviconJiraWordmark1)" d="M61.934 47.414H16.77c0 11.258 9.113 20.371 20.37 20.371h8.31v8.043c0 11.254 9.112 20.367 20.37 20.367V51.301a3.89 3.89 0 0 0-3.886-3.887zm0 0"/><path fill="#253858" d="M46 101.242h3.434v16.98c0 4.481-2.043 7.606-6.786 7.606c-1.78 0-3.175-.297-4.132-.633v-3.254c1.047.422 2.308.633 3.57.633c2.914 0 3.914-1.687 3.914-4.14zm9.742-1.312c1.348 0 2.305.804 2.305 2.238c0 1.394-.957 2.242-2.305 2.242c-1.347 0-2.304-.805-2.304-2.242c0-1.39.957-2.238 2.304-2.238zm-1.695 7.14h3.305v18.59h-3.305zm11.48 18.59H62.31v-18.59h3.218v3.254c1.133-2.199 3.047-3.761 6.785-3.55v3.128c-4.218-.422-6.785.801-6.785 4.774zm22.231-3.34c-1.219 2.453-3.524 3.72-6.485 3.72c-5.085 0-7.652-4.185-7.652-9.677c0-5.238 2.695-9.672 8.047-9.672c2.781 0 4.957 1.223 6.09 3.633v-3.254h3.305v18.59h-3.305zm-5.613.762c2.957 0 5.566-1.816 5.566-5.957v-1.477c0-4.14-2.39-5.957-5.219-5.957c-3.695 0-5.61 2.364-5.61 6.672c.044 4.48 1.872 6.719 5.263 6.719zm0 0"/>`),
		g.Group(children),
	)
}