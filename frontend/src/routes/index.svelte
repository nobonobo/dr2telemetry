<script>
  let Gear = "N";
  let Clutch = 100.0;
  let Brake = 100.0;
  let Throttle = 100.0;
  let Steer = 0;
  runtime.EventsOn("telemetory", (data, deg) => {
    //console.log(data);
    if (data.Gear == 0) {
      Gear = "N";
    } else if (data.Gear < 0) {
      Gear = "R";
    } else {
      Gear = data.Gear.toString(10);
    }
    Clutch = 100 * (1 - data.Clutch);
    Brake = 100 * (1 - data.Brake);
    Throttle = 100 * (1 - data.Throttle);
    Steer = (data.Steer * deg) / 2;
  });
  window.addEventListener("mousemove", () => {
    window.runtime.EventsEmit("window-activate", true);
  });
</script>

<svg
  width="100%"
  height="100%"
  viewBox="0 0 230 120"
  version="1.1"
  id="svg5"
  inkscape:version="1.2 (dc2aedaf03, 2022-05-15)"
  sodipodi:docname="frame.svg"
  xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
  xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
  xmlns="http://www.w3.org/2000/svg"
  xmlns:svg="http://www.w3.org/2000/svg"
>
  <sodipodi:namedview
    id="namedview7"
    pagecolor="#ffffff"
    bordercolor="#000000"
    borderopacity="0.25"
    inkscape:showpageshadow="2"
    inkscape:pageopacity="0.0"
    inkscape:pagecheckerboard="0"
    inkscape:deskcolor="#d1d1d1"
    inkscape:document-units="px"
    showgrid="false"
    showguides="true"
    inkscape:zoom="0.89449008"
    inkscape:cx="187.8165"
    inkscape:cy="183.34468"
    inkscape:window-width="1251"
    inkscape:window-height="637"
    inkscape:window-x="234"
    inkscape:window-y="234"
    inkscape:window-maximized="0"
    inkscape:current-layer="layer1"
    inkscape:lockguides="true"
  />
  <g id="layer1">
    <g id="g995" transform="translate(0,-5)">
      <rect
        style="fill:#d7d71f;fill-opacity:1;stroke:none;stroke-width:0.999934"
        id="Clutch"
        width="20"
        height="100"
        x="20"
        y="15"
      />
      <rect
        style="fill:#407eb6;fill-opacity:1;stroke:none;stroke-width:0.999999"
        id="ClutchInvert"
        width="20"
        height={Clutch}
        x="20"
        y="15"
      />
    </g>
    <g id="g1001" transform="translate(30,-5)">
      <rect
        style="fill:#db1b1b;fill-opacity:1;stroke:none;stroke-width:0.999934"
        id="Brake"
        width="20"
        height="100"
        x="20"
        y="15"
      />
      <rect
        style="fill:#407eb6;fill-opacity:1;stroke:none;stroke-width:1"
        id="BrakeInvert"
        width="20"
        height={Brake}
        x="20"
        y="15"
      />
    </g>
    <g
      id="Steer"
      transform="rotate({Steer},129.9997565,60) translate(-9.9997565)"
    >
      <path
        id="path1055"
        style="fill:#35d422;fill-opacity:1;stroke:none"
        d="m 140,10 a 50,50 0 0 0 -50,50 50,50 0 0 0 50,50 50,50 0 0 0 50,-50 50,50 0 0 0 -50,-50 z m 0,20 a 30,30 0 0 1 30,30 30,30 0 0 1 -30,30 30,30 0 0 1 -30,-30 30,30 0 0 1 30,-30 z"
      />
      <path
        id="circle1057"
        style="fill:#407eb6;fill-opacity:1;stroke:none"
        d="M -11.703125,127.05859 A 50,50 0 0 0 -66.527344,90.427734 50,50 0 0 0 -110,140 a 50,50 0 0 0 43.472656,49.57227 50,50 0 0 0 54.824219,-36.63086 L -31.03125,147.76172 A 30,30 0 0 1 -60,170 a 30,30 0 0 1 -30,-30 30,30 0 0 1 30,-30 30,30 0 0 1 28.867188,22.26562 z"
        transform="rotate(-90)"
      />
    </g>
    <g id="g1846" transform="translate(170,-5)">
      <rect
        style="fill:#0aebec;fill-opacity:1;stroke:none;stroke-width:0.999934"
        id="Throttle"
        width="20"
        height="100"
        x="20"
        y="15"
      />
      <rect
        style="fill:#407eb6;fill-opacity:1;stroke:none;stroke-width:0.999997"
        id="ThrottleInvert"
        width="20"
        height={Throttle}
        x="20"
        y="15"
      />
    </g>
    <text
      xml:space="preserve"
      style="font-size:64px;text-align:center;text-anchor:middle;fill:#2f8bf0;fill-opacity:1;stroke:none;pointer-events:none;"
      x="129.93774"
      y="83.265625"
      id="text3042"
      ><tspan id="Gear" x="129.93774" y="83.265625">{Gear}</tspan></text
    >
  </g>
</svg>
