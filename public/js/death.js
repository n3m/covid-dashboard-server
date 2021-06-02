// Muertos por estado (ambos sexos)
var margin = {
    left:100,
    right: 10,
    top: 10,
    bottom: 100,
};

var width = 600 - margin.left - margin.right,
height = 400 - margin.top - margin.bottom;

var svg = d3.select("#chart-area").append("svg")
    .attr("width", width + margin.right + margin.left)
	.attr("height", height + margin.top + margin.bottom);

var g = svg.append("g").attr("transform", "translate(" + margin.left + ", " + margin.top + ")");
// X Label
g.append("text")
    .attr("class", "x axis-label")
    .attr("x", width/2)
    .attr("y", height+40)
    .attr("font-size", "20px")
    .attr("text-anchor", "middle")
    .style("fill","black")
    .text("States");
// Y Label
g.append("text")
    .attr("class", "y axis-label")
    .attr("x", -(height/2))
    .attr("y", -60)
    .attr("font-size", "20px")
    .attr("text-anchor", "middle")
    .style("fill","black")
    .text("Death Count")
    .attr("transform", "rotate(-90)");

var x = d3.scaleBand().range([0,width]).paddingInner(0.3).paddingOuter(0.3);
var y = d3.scaleLinear().range([height,0]);

var xAxisGroup = g.append("g")    
    .attr("class", "x axis")    
    .attr("transform", "translate(0, " + height + ")")
var yAxisGroup = g.append("g")
    .attr("class", "y axis");

// Graphic starts with both sexs
both();

function both() {
    setValues([]);
}

function male() {
    setValues(
        [
            {
                "eq":"HOMBRE",
            },
        ],
    );
}

function female() {
    setValues(
        [
            {
                "eq":"MUJER",
            },
        ],
    );
}

function setValues(sex) {
    axios.post('/covid', {
        "responseType":"BYSTATE",
        "sexo":sex,
        "defunto":[{
            "gte":0
        }]
    }).then((response)=> {
        update(response);
    }).catch((error) => {
        console.log(error);
    });
}

var t = d3.transition().duration(250);

function update(response) {
    var keys = Object.keys(response.data.data);
    var data = [];
    keys.forEach((key) => {
        data.push(response.data.data[key]);
    });

    var deadStates = data.map((d)=>{
        return d["State"];
    });

    var rectangles = g.selectAll("rect")
        .data(data);

    rectangles.exit()
        .transition(t)
            .attr("y", y(0))
            .attr("height", 0)
        .remove();

    x.domain(deadStates);

    var maxDeadHeight = d3.max(data, (d) =>{
        return d["Count"];
    });
    y.domain([0,maxDeadHeight]);

    var bottomAxis = d3.axisBottom(x);
    xAxisGroup.transition(t).call(bottomAxis);

    var leftAxis = d3.axisLeft(y).ticks(5);
    yAxisGroup.call(leftAxis);

    rectangles.enter().append("rect")
        .attr("y", y(0))    
        .attr("height", 0)
        .attr("x", (d) => { 
            return x(d["State"]); 
        })
        .attr("width", x.bandwidth())
        .merge(rectangles)
        .transition(t)
            .attr("x", (d) => { 
                return x(d["State"]); 
            })
            .attr("width", x.bandwidth())
            .attr("y", (d) => { 
                return y(d["Count"]); 
            } )
            .attr("height", (d) => { 
                return height-y(d["Count"]); 
            })
            .attr("fill", (d) => {
                var fixedColor = d3.rgb(255,0,0);
                fixedColor.r = (d["Count"]/maxDeadHeight)*fixedColor.r;
                return fixedColor;
            });
}