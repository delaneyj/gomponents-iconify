package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunityHealthworkerNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCommunityHealthworkerNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM14.912 13a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm4.235 2c.117 0 .234.014.347.041c1.129.269 1.912.922 2.416 1.86c.44.818.655 1.835.832 2.839l1.058 6a1.5 1.5 0 1 1-2.954.521l-1.059-6c-.124-.702-.235-1.19-.35-1.537L19.146 22c0 1.604 1.796 10.325 2.43 13.351a.99.99 0 0 1-.836 1.186c-4.397.606-7.26.627-11.659.005a.99.99 0 0 1-.83-1.19c.639-2.98 2.425-11.486 2.425-13.352l-.29-3.276c-.116.348-.227.835-.35 1.538l-1.06 6a1.5 1.5 0 0 1-2.954-.521l1.059-6c.177-1.004.391-2.021.832-2.84c.504-.937 1.287-1.59 2.415-1.859a1.5 1.5 0 0 1 .346-.04L12.728 15h6.419Zm-8.272 23.118l.398 3.013a1 1 0 0 0 1.959.125l.766-2.894a40.692 40.692 0 0 1-3.123-.244Zm4.957.267l.76 2.87a1 1 0 0 0 1.958-.124l.386-2.914a38.866 38.866 0 0 1-3.104.168Zm17.58-31.3l7.207 7.208l-1.414 1.414l-1.293-1.293V20h-3v-2.5a1.5 1.5 0 0 0-3 0V20h-3v-5.586l-1.293 1.293l-1.414-1.414l7.207-7.207Zm0 17.5l9.207 9.208l-1.414 1.414l-1.909-1.908v7.74h-3.923v-3.27a1.962 1.962 0 0 0-3.923 0v3.27h-3.923v-7.74l-1.908 1.908l-1.414-1.414l9.207-9.207Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCommunityHealthworkerNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}