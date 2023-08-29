package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserMinusRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="6" r="4"/><path fill-rule="evenodd" d="M13.513 21.487C14.025 22 14.85 22 16.5 22c1.65 0 2.475 0 2.987-.513C20 20.975 20 20.15 20 18.5c0-1.65 0-2.475-.513-2.987C18.975 15 18.15 15 16.5 15c-1.65 0-2.475 0-2.987.513C13 16.025 13 16.85 13 18.5c0 1.65 0 2.475.513 2.987Zm2.404-3.57h-.973a.583.583 0 1 0 0 1.166h3.112a.583.583 0 1 0 0-1.166h-2.139Z" clip-rule="evenodd"/><path d="M18.095 15.031C17.67 15 17.149 15 16.5 15c-1.65 0-2.475 0-2.987.513C13 16.025 13 16.85 13 18.5c0 1.166 0 1.92.181 2.443c-.384.038-.778.057-1.181.057c-3.866 0-7-1.79-7-4s3.134-4 7-4c2.613 0 4.892.818 6.095 2.031Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}