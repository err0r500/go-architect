var colors = d3.scaleOrdinal(d3.schemeCategory10);

var width = window.innerWidth
var height = window.innerHeight
var node
var link;

var projectRootColor = "white"
var projectPackageColor = "#aac9ff"
var corePackageColor = "#8CD28C"
var thirdPartyPackageColor = "#E37474"
var linkInactiveColor = "grey"

var svg = d3.select("body")
    .append("svg")
    .attr("width", "100%")
    .attr("height", "100%")
    .call(d3.zoom().on("zoom", function () {
        svg.attr("transform", d3.event.transform)
    }))
    .append("g")

var simulation = d3.forceSimulation()
    .force("link", d3.forceLink().id(function (d) { return d.id; }).distance(10).strength(0))
    .force("charge", d3.forceManyBody().strength(-1000).distanceMax([50]))
    .force("center", d3.forceCenter(width / 2, height / 2));

d3.json("http://localhost:8080/data/", function (error, graph) {
    if (error) throw error;
    graphSetup(graph.links, graph.nodes);
})


function graphSetup(links, nodes) {
    link = svg.selectAll(".link")
        .data(links)
        .enter()
        .append("line")
        .attr("class", "link")

    link.append("title")
        .text(function (d) { return d.type; });

    node = svg.selectAll(".node")
        .data(nodes)
        .enter()
        .append("g")
        .attr("class", "node")
        .call(d3.drag()
            .on("start", dragstarted)
            .on("drag", dragged)
        );

    node.append("circle")
        .attr("r", (d) => d.label === "projectRoot" ? 15 : 10)
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
        .text(function (d) { return d.name })
        .attr("fill", "grey");

    setupNodesInteractions()

    simulation
        .nodes(nodes)
        .on("tick", ticked);

    simulation.force("link")
        .links(links);
}

function setupNodesInteractions() {
    node.on('mouseover', function (d) {
        var children = []
        var parents = []

        children = buildDepsTree(d.id, "children")
        parents = buildDepsTree(d.id, "parents")

        node.filter((e) => children.indexOf(e.id) !== -1)
            .classed('child', true)
            .classed('active', true);

        node.filter((e) => parents.indexOf(e.id) !== -1)
            .classed('parent', true)
            .classed('active', true);

        link.filter((e) => children.indexOf(e.source.id) !== -1)
            .classed('child', true)
            .classed('active', true);

        link.filter((e) => parents.indexOf(e.target.id) !== -1)
            .classed('parent', true)
            .classed('active', true);
    })

    node.on('mouseout', function (d) {
        link.classed('active', false)
            .classed('child', false)
            .classed('parent', false);
            
        node.classed('active', false)
            .classed('child', false)
            .classed('parent', false);
    })
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
