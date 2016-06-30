var myChart = echarts.init(document.getElementById('main'));

// var ecConfig = require('echarts/config');
// var zrEvent = require('zrender/tool/event');
var curIndx = 0;
var mapType = [
    'china',
    // 23个省
    '广东', '青海', '四川', '海南', '陕西',
    '甘肃', '云南', '湖南', '湖北', '黑龙江',
    '贵州', '山东', '江西', '河南', '河北',
    '山西', '安徽', '福建', '浙江', '江苏',
    '吉林', '辽宁', '台湾',
    // 5个自治区
    '新疆', '广西', '宁夏', '内蒙古', '西藏',
    // 4个直辖市
    '北京', '天津', '上海', '重庆',
    // 2个特别行政区
    '香港', '澳门'
];

var convertData = function (data) {
    var res = [];
    //console.debug("geoMap:", geoCoordMap);
    for (var i = 0; i < data.length; i++) {
        var geoCoord = geoCoordMap[data[i].name];
        if (geoCoord) {
            res.push({
                name: data[i].name,
                value: geoCoord.concat(data[i].value)
            });
        }
    }
    return res;
};


var placeList = [
    {name:'海门', geoCoord:[121.15, 31.89]},

    // {name:'招远', geoCoord:[120.38, 37.35]},

    // {name:'海口', geoCoord:[110.35, 20.02]},

    // {name:'德州', geoCoord:[116.29, 37.45]},

    // {name:'荆州', geoCoord:[112.239741, 30.335165]},

    {name:'大庆', geoCoord:[125.03, 46.58]}
];

var lowData = [
    {name: "沈阳第二营业部", value: 10, geoCoord:[123.487603, 41.806812]},
    {name: "沈阳第三营业部", value: 10, geoCoord:[123.376865, 41.827814]},
    {name: "沈阳第四营业部", value: 10, geoCoord:[123.393993, 41.789829]},
    {name: "沈阳第五营业部", value: 10, geoCoord:[123.371985, 41.79455]},
    {name: "沈阳第十一营业部", value: 10, geoCoord:[123.376865, 41.827814]},
];

var midData = [
    {name:'招远', value: 10, geoCoord:[121.38, 37.35]},
    {name:'海口', value: 10,  geoCoord:[111.35, 20.02]},
    {name:'德州', value: 10, geoCoord:[117.29, 37.45]}

];

var hignData = [
    {name:'招远', value: 10, geoCoord:[120.38, 39.35]},
    {name:'海口', value: 10,  geoCoord:[110.35, 21.02]},
    {name:'德州', value: 10, geoCoord:[116.29, 33.45]}

];

console.debug("lowData:", lowData);

option = {
    // backgroundColor: '#1b1b1b',
    color: [
        'rgba(255, 0, 0, 1)',
        'rgba(1, 255, 0, 1)',
        'rgba(1, 1, 255, 1)'
    ],
    title : {
        text: '大规模MarkPoint特效',
        subtext: '纯属虚构',
        x:'center',
        // textStyle : {
        //     color: '#fff'
        // }
    },
    legend: {
        orient: 'vertical',
        x:'left',
        data:['强','中','弱'],
        // textStyle : {
        //     color: '#fff'
        // }
    },
    // toolbox: {
    //     show : true,
    //     orient : 'vertical',
    //     x: 'right',
    //     y: 'center',
    //     feature : {
    //         mark : {show: true},
    //         dataView : {show: true, readOnly: false},
    //         restore : {show: true},
    //         saveAsImage : {show: true}
    //     }
    // },
    series : [
        {
            name: '弱',
            type: 'map',
            mapType: 'china',
            selectedMode : 'single',
            data : lowData,
            markPoint : {

                symbol: 'diamond',
                symbolSize: 6,
                large: true,
                // effect : {
                //     show: true
                // },
                itemStyle:{
                    normal:{
                        label:{show:true},
                        borderColor:'rgba(100,149,237,0.8)',
                        borderWidth:1.5,
                    },
                    // emphasis:{label:{show:true}}
                    emphasis: {
                        borderColor: '#1e90ff',
                        borderWidth: 5,
                        label: {
                            show: true
                        }
                    }
                },
                data: lowData
            }
        },
        {
            name: '中',
            type: 'map',
            mapType: 'china',
            data : [],
            itemStyle:{
                normal:{
                    label:{show:true},
                    borderColor:'rgba(100,149,237,0.8)',
                    borderWidth:1.5,
                },
                emphasis:{label:{show:true}}
            },
            markPoint : {
                symbolSize: 3,
                large: true,
                effect : {
                    show: true
                },
                data: midData
            }
        },
        {
            name: '强',
            type: 'map',
            mapType: 'china',
            hoverable: false,
            // roam:true,
            data : [],
            itemStyle:{
                normal:{
                    label:{show:true},
                    borderColor:'rgba(100,149,237,0.8)',
                    borderWidth:1.5,
                },
                emphasis:{label:{show:true}}
            },
            markPoint : {
                symbol : 'diamond',
                symbolSize: 6,
                large: true,
                effect : {
                    show: true
                },
                data: hignData
            }
        }
    ]
}

myChart.on(echarts.config.EVENT.MAP_SELECTED, function (param){
    console.debug("click event fired.");

    var len = mapType.length;
    var mt = mapType[curIndx % len];
    if (mt == 'china') {
        // 全国选择时指定到选中的省份
        var selected = param.selected;
        for (var i in selected) {
            if (selected[i]) {
                mt = i;
                while (len--) {
                    if (mapType[len] == mt) {
                        curIndx = len;
                    }
                }
                break;
            }
        }
        // option.tooltip.formatter = '点击返回全国<br/>{b}';
    }
    else {
        curIndx = 0;
        mt = 'china';
        // option.tooltip.formatter = '点击进入该省<br/>{b}';
    }
    option.series[0].mapType = mt;
    option.series[1].mapType = mt;
    option.series[2].mapType = mt;
    option.title.subtext = mt + ' （点击切换）';
    myChart.setOption(option, true);
});

myChart.setOption(option);
