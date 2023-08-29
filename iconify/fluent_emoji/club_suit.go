package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClubSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f288id1)"><path fill="url(#f288id0)" d="M23.767 23.426c3.448-.17 6.243-2.941 6.375-6.323c.153-3.772-2.917-6.874-6.722-6.874h-.01a.596.596 0 0 1-.602-.68c.062-.41.082-.83.062-1.27c-.174-3.403-3-6.144-6.467-6.274c-3.835-.14-6.987 2.871-6.987 6.593c0 .32.02.64.071.95c.051.361-.245.661-.602.661h-.01c-3.805 0-6.875 3.102-6.722 6.874c.133 3.392 2.928 6.153 6.375 6.323c2 .1 3.825-.66 5.12-1.94c.174-.17.46-.05.46.18v1.26c0 2.101-1.714 3.812-3.856 3.852h-.03c-.908 0-1.694.7-1.714 1.591c-.02.91.734 1.651 1.652 1.651h11.944c.715 0 1.367-.44 1.571-1.1c.357-1.121-.49-2.142-1.57-2.142h-.082c-2.132-.04-3.856-1.75-3.856-3.852v-1.26c0-.24.296-.35.46-.18a6.804 6.804 0 0 0 5.14 1.96Z"/></g><defs><linearGradient id="f288id0" x1="24.465" x2="11.602" y1="11.078" y2="29.196" gradientUnits="userSpaceOnUse"><stop stop-color="#534165"/><stop offset="1" stop-color="#3F3946"/></linearGradient><filter id="f288id1" width="29.75" height="29.5" x="1.148" y="1.25" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.427451 0 0 0 0 0.372549 0 0 0 0 0.482353 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18_4053"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75" dy="-.75"/><feGaussianBlur stdDeviation=".625"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.172549 0 0 0 0 0.109804 0 0 0 0 0.227451 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18_4053" result="effect2_innerShadow_18_4053"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75" dy=".75"/><feGaussianBlur stdDeviation=".625"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.196078 0 0 0 0 0.192157 0 0 0 0 0.2 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18_4053" result="effect3_innerShadow_18_4053"/></filter></defs></g>`),
		g.Group(children),
	)
}