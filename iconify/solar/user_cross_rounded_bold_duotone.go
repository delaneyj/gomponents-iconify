package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCrossRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="6" r="4"/><path fill-rule="evenodd" d="M16.5 22c-1.65 0-2.475 0-2.987-.513C13 20.975 13 20.15 13 18.5c0-1.65 0-2.475.513-2.987C14.025 15 14.85 15 16.5 15c1.65 0 2.475 0 2.987.513C20 16.025 20 16.85 20 18.5c0 1.65 0 2.475-.513 2.987C18.975 22 18.15 22 16.5 22Zm-1.143-5.468a.583.583 0 1 0-.825.825l1.143 1.143l-1.143 1.143a.583.583 0 1 0 .825.825l1.143-1.143l1.143 1.143a.583.583 0 1 0 .825-.825L17.325 18.5l1.143-1.143a.583.583 0 1 0-.825-.825L16.5 17.675l-1.143-1.143Z" clip-rule="evenodd"/><path d="M18.095 15.031C17.67 15 17.149 15 16.5 15c-1.65 0-2.475 0-2.987.513C13 16.025 13 16.85 13 18.5c0 1.166 0 1.92.181 2.443c-.384.038-.778.057-1.181.057c-3.866 0-7-1.79-7-4s3.134-4 7-4c2.613 0 4.892.818 6.095 2.031Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}