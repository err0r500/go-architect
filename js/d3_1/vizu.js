var colors = d3.scaleOrdinal(d3.schemeCategory10);

var width = window.innerWidth
var height = window.innerHeight
var node
var link;

var projectRootColor = "white"
var projectPackageColor = "#aac9ff"
var corePackageColor = "#8CD28C"
var thirdPartyPackageColor = "#E37474"
var currDependsOnColor = "#BA4343"
var dependsOnCurrColor = "#359535"
var linkInactiveColor = "grey"

var svg = d3.select("body")
    .append("svg")
    .attr("width", "100%")
    .attr("height", "100%")
    .call(d3.zoom().on("zoom", function () {
        svg.attr("transform", d3.event.transform)
    }))
    .append("g")

svg.append('defs').append('marker')
    .attrs({
        'id': 'arrowhead',
        'viewBox': '-0 -5 10 20',
        'refX': 13,
        'refY': 0,
        'orient': 'auto',
        'markerWidth': 5,
        'markerHeight': 10,
        'xoverflow': 'visible'
    })
    .append('svg:path')
    .attr('d', 'M 0,-5 L 10 ,0 L 0,5')
    .attr('fill', '#999')
    .style('stroke', 'none');

var simulation = d3.forceSimulation()
    .force("link", d3.forceLink().id(function (d) { return d.id; }).distance(10).strength(0))
    .force("charge", d3.forceManyBody().strength(-1000).distanceMax([50]))
    .force("center", d3.forceCenter(width / 2, height / 2));

d3.json("../testGraph.json", function (error, graph) {
    if (error) throw error;
    graphSetup(graph.links, graph.nodes);
})

function getLinkClassFromCoupling(l) {
    cssClass = "link"
    if (l.Type == "low") {
        cssClass += " lowCoupling"
    }
    return cssClass
}

function graphSetup(links, nodes) {
    link = svg.selectAll(".link")
        .data(links)
        .enter()
        .append("line")
        .attr("class", (d) => getLinkClassFromCoupling(d) )
        .attr('marker-end', 'url(#arrowhead)')

    link.append("title")
        .text(function (d) { return d.type; });

    edgepaths = svg.selectAll(".edgepath")
        .data(links)
        .enter()
        .append('path')
        .attrs({
            'class': 'edgepath',
            'fill-opacity': 0,
            'stroke-opacity': 0,
            'id': function (d, i) { return 'edgepath' + i }
        })
        .style("pointer-events", "none");

    edgelabels = svg.selectAll(".edgelabel")
        .data(links)
        .enter()
        .append('text')
        .style("pointer-events", "none")
        .attrs({
            'class': 'edgelabel',
            'id': function (d, i) { return 'edgelabel' + i },
            'font-size': 10,
            'fill': 'white'
        });

    edgelabels.append('textPath')
        .attr('xlink:href', function (d, i) { return '#edgepath' + i })
        .style("text-anchor", "middle")
        .style("pointer-events", "none")
        .attr("startOffset", "50%")
        .text(function (d) { return d.type });

    node = svg.selectAll(".node")
        .data(nodes)
        .enter()
        .append("g")
        .call(d3.drag()
            .on("start", dragstarted)
            .on("drag", dragged)
        );

    node.on('mouseover', function (d) {
        d3.select(this).attr("class", "active")

        link.filter((e) => buildDepsTree(d.id, "children").indexOf(e.source.id) !== -1)
        .style("stroke", currDependsOnColor)
        .classed('active', true);
        
        link.filter((e) => buildDepsTree(d.id, "parents").indexOf(e.target.id) !== -1)
        .style("stroke", dependsOnCurrColor)
        .classed('active', true);
    })
    
    node.on('mouseout', function (d) {
        link.classed('active', false)
    })

    node.append("circle")
        .attr("r",(d) => d.label === "projectRoot" ? 15 : 10)
        .style("fill", function (d, i) {
            switch (d.label) {
                case "projectRoot":
                    return projectRootColor
                case "projectPackage":
                    return projectPackageColor
                case "corePackage":
                    return corePackageColor
                default:
                    return thirdPartyPackageColor
            }
        })

    node.append("text")
        .attr("dy", -13)
        .attr("dx", -13)
        .style("pointer-events", "none")        
        .text(function (d) { return d.name }).attr("fill", "grey")

    simulation
        .nodes(nodes)
        .on("tick", ticked);

    simulation.force("link")
        .links(links);
}

function buildDepsTree(startNodeID, searchForParentsOrChildren) {
    var srcArray = [startNodeID]
    
    var from = "source"
    var to = "target"
    if (searchForParentsOrChildren === "parents") {
        from = "target"
        to = "source"
    }

    for (i = 0; i < srcArray.length; i++) {
        link.each((e) => {
            if (srcArray[i] === e[from].id && srcArray.indexOf(e[to].id) === -1) {
                srcArray.push(e[to].id)
            }
        })
    }
    return srcArray
}

function ticked() {
    link
        .attr("x1", function (d) { return d.source.x; })
        .attr("y1", function (d) { return d.source.y; })
        .attr("x2", function (d) { return d.target.x; })
        .attr("y2", function (d) { return d.target.y; });

    node
        .attr("transform", function (d) { return "translate(" + d.x + ", " + d.y + ")"; });

    edgepaths.attr('d', function (d) {
        return 'M ' + d.source.x + ' ' + d.source.y + ' L ' + d.target.x + ' ' + d.target.y;
    });

    edgelabels.attr('transform', function (d) {
        if (d.target.x < d.source.x) {
            var bbox = this.getBBox();

            rx = bbox.x + bbox.width / 2;
            ry = bbox.y + bbox.height / 2;
            return 'rotate(180 ' + rx + ' ' + ry + ')';
        }
        else {
            return 'rotate(0)';
        }
    });
}

function dragstarted(d) {
    if (!d3.event.active) simulation.alphaTarget(0.3).restart()
    d.fx = d.x;
    d.fy = d.y;
}

function dragged(d) {
    d.fx = d3.event.x;
    d.fy = d3.event.y;
}
