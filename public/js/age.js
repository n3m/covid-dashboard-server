// Infectados por edad
var margin = {
    left:50,
    right: 10,
    top: 10,
    bottom: 100,
};

var width = 1400 - margin.left - margin.right,
height = 600 - margin.top - margin.bottom;

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
    .text("Ages");
// Y Label
g.append("text")
    .attr("class", "y axis-label")
    .attr("x", -(height/2))
    .attr("y", -60)
    .attr("font-size", "20px")
    .attr("text-anchor", "middle")
    .style("fill","black")
    .text("Infected Count")
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
    setValues([],"black");
}

function male() {
    setValues(
        [
            {
                "eq":"HOMBRE",
            },
        ],
        "cyan",
    );
}

function female() {
    setValues(
        [
            {
                "eq":"MUJER",
            },
        ],
        "pink",
    );
}

function setValues(sex, color) {
    axios.post('/covid', {
        "responseType":"BYAGE",
        "sexo":sex,
    }).then((response)=> {
        update(response, color);
    }).catch((error) => {
        console.log(error);
    });
}

var t = d3.transition().duration(250);

function update(response, color) {
    // Get count by ages
    var agesCount = response.data.data["AgesByCount"];
    var ages = Object.keys(agesCount);
    var data = [];
    ages.forEach((key) => {
        data.push({
            age:key, 
            count:agesCount[key]
        });
    });

    var rectangles = g.selectAll("rect")
        .data(data);

    rectangles.exit()
        .transition(t)
            .attr("y", y(0))
            .attr("height", 0)
        .remove();

    x.domain(ages);

    var maxCountHeight = d3.max(data, (d) =>{
        return d.count;
    });
    y.domain([0,maxCountHeight]);

    var bottomAxis = d3.axisBottom(x);
    xAxisGroup.transition(t).call(bottomAxis);
    
    var leftAxis = d3.axisLeft(y).ticks(5);
    yAxisGroup.call(leftAxis);

    rectangles.enter().append("rect")
        .attr("y", y(0))    
        .attr("height", 0)
        .attr("x", (d) => { 
            return x(d.age); 
        })
        .attr("width", x.bandwidth())
        .merge(rectangles)
        .transition(t)
            .attr("x", (d) => { 
                return x(d.age); 
            })
            .attr("width", x.bandwidth())
            .attr("y", (d) => { 
                return y(d.count); 
            } )
            .attr("height", (d) => { 
                return height-y(d.count); 
            })
            .attr("fill", color);  
}