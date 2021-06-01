// Origen
var margin = {
    left:100,
    right: 10,
    top: 10,
    bottom: 100,
};

var width = 600;
var height = 400;

var svg = d3.select("#chart-area").append("svg")
    .attr("width", width + margin.right + margin.left)
	.attr("height", height + margin.top + margin.bottom);

var g = svg.append("g")
.attr("transform", "translate(" + margin.left + ", " + margin.top + ")");

axios.post('/covid', {
    "responseType":"BYPRIVPUB",
}).then((response) => {
    console.log(response.data.data);
    var data = [{
        name:"Private",
        count: response.data.data["PrivCount"],
    }, {
        name:"Public",
        count: response.data.data["PubCount"],
    }];

    var origins = data.map((d) => {
        return d.name;
    });

    var x = d3.scaleBand()
    .domain(origins)
    .range([0,width])
    .paddingInner(0.3)
    .paddingOuter(0.3);

    var maxCountHeight = d3.max(data, (d) =>{
        return d.count;
    });
    var y = d3.scaleLinear()
    .domain([0,maxCountHeight])
    .range([height,0]);

    var rectangles = g.selectAll("rect")
    .data(data);

    var bottomAxis = d3.axisBottom(x);
    g.append("g")    
        .attr("class", "bottom axis")    
        .attr("transform", "translate(0, " + height + ")")    
        .call(bottomAxis)        
    .selectAll("text")
        .attr("y", "10")
        .attr("x", "10")
        .attr("text-anchor", "end")
    // X Label
    g.append("text")
    .attr("class", "x axis-label")
    .attr("x", width/2)
    .attr("y", height+40)
    .attr("font-size", "20px")
    .attr("text-anchor", "middle")
    .style("fill","black")
    .text("Sector");

    var leftAxis = d3.axisLeft(y).ticks(5)
    g.append("g")
        .attr("class", "left axis")
        .call(leftAxis);
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

    rectangles.enter()
    .append("rect")
    .attr("x", (d) => { 
        return x(d.name); 
    })
    .attr("y", (d) => { 
        return y(d.count); 
    } )
    .attr("width", x.bandwidth())
    .attr("height", (d) => { 
        return height-y(d.count); 
    })
    .attr("fill", "cyan");

}).catch((error) => {
    console.log(error);
});